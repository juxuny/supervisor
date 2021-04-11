package proxy

import (
	"fmt"
	"io"
	"net"
	"runtime/debug"
	"sync"
	"time"
)

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
					fmt.Println(err)
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
	return nil
}

func (t *Server) start() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", t.proxy.ListenPort))
	if err != nil {
		fmt.Println(err)
		return
	}
	failedCount := 0
	for {
		client, err := ln.Accept()
		if err != nil {
			failedCount += 1
			if failedCount > 5 {
				break
			}
			continue
		}
		go t.serveClient(client)
		failedCount = 0
	}
}

func (t *Server) serveClient(conn net.Conn) {
	remoteConn, err := net.Dial("tcp", t.proxy.Remote)
	if err != nil {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
		return
	}
	go func() {
		_, err := io.Copy(conn, remoteConn)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	_, err = io.Copy(remoteConn, conn)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Server) Status() (ret *Status, err error) {
	ret = &Status{}
	ret.ListenPort = t.proxy.ListenPort
	ret.Remote = t.proxy.Remote
	return
}
