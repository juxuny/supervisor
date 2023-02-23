package proxy

import (
	"context"
	"fmt"
	"github.com/juxuny/env"
	"github.com/juxuny/supervisor/log"
	"github.com/juxuny/supervisor/trace"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"path"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

const BlockSize = 10 * (1 << 10) // 1k

type IServer interface {
	Start()
	UpdateRemote(remote string) error
	Status() (*Status, error)
}

type Server struct {
	proxy Proxy
	sync.Mutex

	ln net.Listener
}

func NewServer(proxy Proxy) IServer {
	s := &Server{proxy: proxy}
	return s
}

func (t *Server) Start() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error(err)
					debug.PrintStack()
				}
			}()
			t.start()
		}()
		time.Sleep(time.Second * 3)
	}
}

func (t *Server) UpdateRemote(remote string) error {
	t.Lock()
	defer t.Unlock()
	t.proxy.Remote = remote
	if err := SaveRemote(int(t.proxy.ControlPort), remote); err != nil {
		log.Error(err)
	}
	return nil
}

func (t *Server) start() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", t.proxy.ListenPort))
	if err != nil {
		log.Error(err)
		return
	}
	failedCount := 0
	trace.InitReqId("start")
	for {
		client, err := ln.Accept()
		if err != nil {
			failedCount += 1
			if failedCount > 5 {
				log.Error("failed count:", failedCount)
				break
			}
			log.Error(err)
			continue
		}
		log.Info("accepted:", client.RemoteAddr().String())
		trace.GoRun(func() {
			trace.InitReqId()
			t.serveClient(client)
		})
		failedCount = 0
	}
}

func (t *Server) transfer(ctx context.Context, cancel context.CancelFunc, from net.Conn, to net.Conn) {
	buf := make([]byte, BlockSize)
	defer func() {
		_ = from.Close()
		_ = to.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if t.proxy.ReadTimeout > 0 {
			if err := from.SetDeadline(time.Now().Add(time.Second * time.Duration(t.proxy.ReadTimeout))); err != nil {
				log.Error(err)
			}
			if err := to.SetDeadline(time.Now().Add(time.Second * time.Duration(t.proxy.ReadTimeout))); err != nil {
				log.Error(err)
			}
		}
		n, err := from.Read(buf)
		if err != nil {
			log.Error(err)
			cancel()
			return
		}
		_, err = to.Write(buf[:n])
		if err != nil {
			log.Error(err)
			cancel()
			return
		}
	}
}

func (t *Server) serveClient(conn net.Conn) {
	remoteConn, err := net.Dial("tcp", t.proxy.Remote)
	if err != nil {
		log.Info("connect to backend failed:", err)
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
		return
	}
	log.Info("connected to backend:", t.proxy.Remote)
	//go func() {
	//	buf := make([]byte, BlockSize)
	//	for {
	//		_ = conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	//		n, err := conn.Read(buf)
	//		if err != nil {
	//			_ = remoteConn.Close()
	//			return
	//		}
	//		_, err = remoteConn.Write(buf[:n])
	//		if err != nil {
	//			_ = conn.Close()
	//			return
	//		}
	//	}
	//}()
	//buf := make([]byte, BlockSize)
	//for {
	//	n, err := remoteConn.Read(buf)
	//	if err != nil {
	//		_ = conn.Close()
	//		return
	//	}
	//	_, err = conn.Write(buf[:n])
	//	if err != nil {
	//		_ = remoteConn.Close()
	//		return
	//	}
	//}
	ctx, cancel := context.WithCancel(context.Background())
	trace.GoRun(func() {
		log.Info("start pass to_backend, remote addr: ", remoteConn.RemoteAddr().String())
		t.transfer(ctx, cancel, conn, remoteConn)
		log.Info("finished to_backend")
	})
	trace.GoRun(func() {
		log.Info("start pass to_client, remote addr: ", conn.RemoteAddr().String())
		t.transfer(ctx, cancel, remoteConn, conn)
		log.Info("finished to_client")
	})
	<-ctx.Done()
	//go t.transfer(ctx, cancel, conn, remoteConn)
	//go t.transfer(ctx, cancel, remoteConn, conn)
}

func (t *Server) Status() (ret *Status, err error) {
	ret = &Status{}
	ret.ListenPort = t.proxy.ListenPort
	ret.Remote = t.proxy.Remote
	return
}

var fileLock = &sync.Mutex{}
var dataDir = env.GetString("DATA_DIR", "/data")

func SaveRemote(controlPort int, remote string) error {
	fileLock.Lock()
	defer fileLock.Unlock()
	fileName := path.Join(dataDir, fmt.Sprintf("%d.remote", controlPort))
	if err := ioutil.WriteFile(fileName, []byte(remote), 0666); err != nil {
		return errors.Wrap(err, "save remote failed")
	}
	return nil
}

func GetRemoteFromFile(controlPort int) (remote string, err error) {
	fileLock.Lock()
	defer fileLock.Unlock()
	fileName := path.Join(dataDir, fmt.Sprintf("%d.remote", controlPort))
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", errors.Wrap(err, "read remote config failed")
	}
	return strings.Trim(string(data), " \r\n\t"), nil
}
