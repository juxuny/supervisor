package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type server struct {
	supervisor.UnimplementedSupervisorServer
}

func getProxyClient(host string) (client proxy.ProxyClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = proxy.NewProxyClient(conn)
	return client, nil
}

func getSupervisorClient(host string) (client supervisor.SupervisorClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = supervisor.NewSupervisorClient(conn)
	return client, nil
}

func (t *server) ProxyStatus(ctx context.Context, req *proxy.StatusReq) (resp *proxy.StatusResp, err error) {
	return resp, nil
}
func (t *server) Apply(ctx context.Context, req *supervisor.ApplyReq) (resp *supervisor.ApplyResp, err error) {
	resp = &supervisor.ApplyResp{}
	dockerClient, err := supervisor.NewDockerClient(supervisor.NewDefaultDockerClientConfig())
	if err != nil {
		return nil, err
	}
	_, err = dockerClient.Apply(ctx, *req.Config)
	if err != nil {
		return nil, err
	}
	resp.Msg = "success"
	return resp, nil
}

func (t *server) Get(ctx context.Context, req *supervisor.GetReq) (resp *supervisor.GetResp, err error) {
	return
}
