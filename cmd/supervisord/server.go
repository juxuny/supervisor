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

func (t *server) ProxyStatus(ctx context.Context, req *proxy.StatusReq) (resp *proxy.StatusResp, err error) {
	return resp, nil
}
func (t *server) Apply(ctx context.Context, req *supervisor.ApplyReq) (resp *supervisor.ApplyResp, err error) {
	return
}

func (t *server) Get(ctx context.Context, req *supervisor.GetReq) (resp *supervisor.GetResp, err error) {
	return
}
