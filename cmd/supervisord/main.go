package main

import (
	"flag"
	"fmt"
	"github.com/juxuny/supervisor"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	configFile string
)

type server struct {
	supervisor.UnimplementedSupervisorServer
}

func main() {
	flag.StringVar(&configFile, "c", "supervisor.yaml", "config file")
	flag.Parse()
	config, err := supervisor.Parse(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	addr := fmt.Sprintf(":%d", config.Supervisor.ControlPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("listen", addr)
	supervisor.RegisterSupervisorServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
