package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedProxyServer
}

func (s *server) Status(ctx context.Context, in *pb.StatusReq) (*pb.StatusResp, error) {
	var resp pb.StatusResp
	currentStatus, _ := proxyServer.Status()
	resp.Status = &pb.Status{
		ListenPort: currentStatus.ListenPort,
		Remote:     currentStatus.Remote,
	}
	return &resp, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateReq) (*pb.UpdateResp, error) {
	var resp pb.UpdateResp
	if err := proxyServer.UpdateRemote(in.Status.Remote); err != nil {
		return nil, errors.Wrap(err, "update failed")
	}
	currentStatus, _ := proxyServer.Status()
	resp.Status = &pb.Status{
		ListenPort: currentStatus.ListenPort,
		Remote:     currentStatus.Remote,
	}
	return &resp, nil
}

var (
	configFile  string
	proxyServer pb.IServer
)

func main() {
	flag.StringVar(&configFile, "c", "config/proxy.yaml", "proxy config yaml")
	flag.Parse()
	proxyConfig, err := pb.Parse(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	proxyServer = pb.NewServer(proxyConfig.Proxy)
	go proxyServer.Start()
	addr := fmt.Sprintf(":%d", int(proxyConfig.Proxy.ControlPort))
	fmt.Println("listen ", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
