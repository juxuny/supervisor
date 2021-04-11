package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/juxuny/supervisor/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedProxyServer
}

func (s *server) Status(ctx context.Context, in *pb.StatusReq) (*pb.StatusResp, error) {
	return &pb.StatusResp{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}

var (
	configFile string
)

func main() {
	flag.StringVar(&configFile, "c", "config/proxy.yaml", "proxy config yaml")
	flag.Parse()
	proxyConfig, err := pb.Parse(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	proxyServer := pb.NewServer(proxyConfig.Proxy)
	proxyServer.Start()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", int(proxyConfig.Proxy.ControlPort)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
