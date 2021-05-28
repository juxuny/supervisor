package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/spf13/cobra"
	"os"
)

var updateFlag = struct {
	supervisor.BaseFlag
	Remote string
}{}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update remote address",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update remote address")
		ctx, cancel := context.WithTimeout(context.Background(), Timeout)
		defer cancel()
		if len(updateFlag.Host) == 0 {
			fmt.Println("missing --host")
			os.Exit(-1)
		}
		client, err := getClient(statusFlag.Host[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		_, err = client.Update(ctx, &pb.UpdateReq{
			Status: &pb.Status{
				Remote: updateFlag.Remote,
			},
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(255)
		}
		checkStatus()
	},
}

func init() {
	updateCmd.PersistentFlags().StringSliceVar(&updateFlag.Host, "host", []string{fmt.Sprintf("127.0.0.1:%d", DefaultControlPlanePort)}, "host")
	updateCmd.PersistentFlags().StringVar(&updateFlag.Remote, "remote", "", "remote address")
	rootCmd.AddCommand(updateCmd)
}
