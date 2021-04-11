package main

import (
	"fmt"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	DefaultControlPlanePort = 50050
	Timeout                 = time.Second * 15
)

var (
	rootCmd = &cobra.Command{
		Use:   "proxy-ctl",
		Short: "proxy-ctl",
	}
)

func getClient(host string) (client pb.ProxyClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = pb.NewProxyClient(conn)
	return client, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
