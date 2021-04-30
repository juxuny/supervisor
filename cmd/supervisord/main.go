package main

import (
	"flag"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/juxuny/supervisor/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	configFile   string
	globalConfig supervisor.Config
	logger       = log.NewLogger("[sup]")
)

func main() {
	flag.StringVar(&configFile, "c", "supervisor.yaml", "config file")
	flag.Parse()
	config, err := supervisor.Parse(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	globalConfig = config.Supervisor
	if err := supervisor.Init(globalConfig); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	addr := fmt.Sprintf(":%d", config.Supervisor.ControlPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen:", err)
	}
	defer func() {
		if err := ln.Close(); err != nil {
			logger.Error(err)
		}
	}()
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		logger.Error("cannot load TLS credentials: ", err)
		os.Exit(-1)
	}
	s := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	fmt.Println("listen", addr)
	supervisor.RegisterSupervisorServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		logger.Error("failed to serve:", err)
	}
}
