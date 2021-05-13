package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/juxuny/supervisor/proxy"
	"github.com/spf13/cobra"
	"os"
)

var checkFlag = struct {
	supervisor.BaseFlag
	Remote string
	Type   int
	Port   int
	Path   string
}{}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "get check",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health check")
		ctx, cancel := context.WithTimeout(context.Background(), Timeout)
		defer cancel()
		client, err := getClient(statusFlag.Host)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		_, err = client.Check(ctx, &proxy.CheckReq{
			Type: proxy.HealthCheckType(checkFlag.Type),
			Host: checkFlag.Remote,
			Path: checkFlag.Path,
			Port: uint32(checkFlag.Port),
		})
		if err != nil {
			fmt.Println("failed:", err)
			os.Exit(-1)
		}
		fmt.Println("health check success")
	},
}

func init() {
	checkCmd.PersistentFlags().IntVar(&checkFlag.Port, "port", 8080, "port")
	checkCmd.PersistentFlags().StringVar(&checkFlag.Path, "path", "/healthz", "health check path")
	checkCmd.PersistentFlags().StringVar(&checkFlag.Remote, "remote", "127.0.0.1", "service address")
	checkCmd.PersistentFlags().IntVar(&checkFlag.Type, "type", 0, "0=http, 1=tcp, default=0")
	checkCmd.PersistentFlags().StringVar(&checkFlag.Host, "host", "127.0.0.1:50050", "host")
	rootCmd.AddCommand(checkCmd)
}
