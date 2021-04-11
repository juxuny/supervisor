package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/spf13/cobra"
	"os"
)

var statusFlag = struct {
	supervisor.BaseFlag
}{}

func checkStatus() {
	fmt.Println("check status...")
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	client, err := getClient(statusFlag.Host)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	resp, err := client.Status(ctx, &pb.StatusReq{})
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
	fmt.Println("listen port:", resp.Status.ListenPort)
	fmt.Println("remote:", resp.Status.Remote)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "get status",
	Run: func(cmd *cobra.Command, args []string) {
		checkStatus()
	},
}

func init() {
	statusCmd.PersistentFlags().StringVar(&statusFlag.Host, "host", "127.0.0.1:50050", "host")
	rootCmd.AddCommand(statusCmd)
}
