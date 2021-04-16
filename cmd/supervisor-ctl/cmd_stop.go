package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"os"
)

var stopFlag = struct {
	supervisor.BaseFlag
	Name string
}{}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), supervisor.DefaultTimeout)
		defer cancel()
		client, err := getClient(ctx, stopFlag.Host)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		_, err = client.Stop(ctx, &supervisor.StopReq{Name: stopFlag.Name})
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
	},
}

func init() {
	stopCmd.PersistentFlags().StringVar(&stopFlag.Host, "host", "127.0.0.1:50060", "host")
	stopCmd.PersistentFlags().StringVar(&stopFlag.Name, "name", "", "service name")
	rootCmd.AddCommand(stopCmd)
}
