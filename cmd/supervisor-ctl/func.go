package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"strconv"
	"strings"
)

func getClient(ctx context.Context, host string) (client supervisor.SupervisorClient, err error) {
	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = supervisor.NewSupervisorClient(conn)
	return client, nil
}

func parseBlockSize(s string) (blockSize int64, err error) {
	s = strings.ToLower(s)
	base := int64(1)
	if strings.Contains(s, "k") {
		base *= 1 << 10
		s = strings.Replace(s, "k", "", 1)
	} else if strings.Contains(s, "m") {
		base *= 1 << 20
		s = strings.Replace(s, "m", "", 1)
	} else if strings.Contains(s, "g") {
		base *= 1 << 30
		s = strings.Replace(s, "g", "", 1)
	} else if strings.Contains(s, "t") {
		base *= 1 << 40
		s = strings.Replace(s, "t", "", 1)
	}
	blockSize, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return blockSize, err
	}
	blockSize *= base
	return
}
