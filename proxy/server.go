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
	Start() error
	UpdateRemote(remote string) error
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

func (t *Server) Start() error {
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
