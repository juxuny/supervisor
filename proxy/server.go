package proxy

import (
	"fmt"
	"github.com/juxuny/env"
	"github.com/juxuny/supervisor/log"
	"github.com/juxuny/supervisor/trace"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
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
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error(err)
					debug.PrintStack()
				}
			}()
			if t.proxy.Http {
				t.startHttpProxy()
			} else {
				t.start()
			}
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

func (t *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestPath := fmt.Sprintf("http://%s%s", t.proxy.Remote, r.RequestURI)
	req, err := http.NewRequest(r.Method, requestPath, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Error("http error,", http.StatusBadGateway, r.RequestURI)
		return
	}
	delHopHeaders(r.Header)
	copyHeader(req.Header, r.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		log.Error("http error", http.StatusBadGateway, req.RequestURI)
		return
	}
	defer resp.Body.Close()
	delHopHeaders(resp.Header)
	log.Debug(req.URL.String())
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	_, err = w.Write(buf)
	if err != nil {
		log.Error(err)
	}
}

func (t *Server) startHttpProxy() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", t.proxy.ListenPort), t); err != nil {
		log.Error(err)
	}
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

func (t *Server) transfer(from net.Conn, to net.Conn) {
	defer func() {
		_ = from.Close()
		_ = to.Close()
	}()
	buf := make([]byte, BlockSize)
	for {
		n, err := from.Read(buf)
		if err != nil {
			break
		}
		_, err = to.Write(buf[:n])
		if err != nil {
			break
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
	go t.transfer(remoteConn, conn)
	go t.transfer(conn, remoteConn)
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
