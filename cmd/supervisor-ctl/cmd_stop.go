package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var stopFlag = struct {
	Name string
}{}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop",
	Run: func(cmd *cobra.Command, args []string) {
		if stopFlag.Name == "" {
			Fatal("name cannot be empty")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(baseFlag.Timeout)*time.Second)
		defer cancel()
		client, err := getClient(ctx, baseFlag.Host, baseFlag.CertFile)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		_, err = client.Stop(ctx, &supervisor.StopReq{Name: stopFlag.Name})
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		logger.Println("stop ", stopFlag.Name, " success")
	},
}

func init() {
	initBaseFlag(stopCmd)
	stopCmd.PersistentFlags().StringVar(&stopFlag.Name, "name", "", "service name")
	rootCmd.AddCommand(stopCmd)
}
