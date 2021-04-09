package main

import (
	"context"
	"log"
	"net"

	pb "github.com/juxuny/supervisor/proxy"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedProxyServer
}

func (s *server) Status(ctx context.Context, in *pb.ProxyStatusReq) (*pb.ProxyStatusResp, error) {
	return &pb.ProxyStatusResp{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.ProxyUpdateReq) (*pb.ProxyUpdateResp, error) {
	return &pb.ProxyUpdateResp{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
