package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func getClient(ctx context.Context, host string) (client supervisor.SupervisorClient, err error) {
	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = supervisor.NewSupervisorClient(conn)
	return client, nil
}
