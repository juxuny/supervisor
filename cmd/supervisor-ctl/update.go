package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"os"
)

var applyFlag = struct {
	supervisor.BaseFlag
	File string
}{}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply",
	Run: func(cmd *cobra.Command, args []string) {
		if applyFlag.File == "" {
			logger.Error("missing file")
			os.Exit(-1)
		}
		dc, err := supervisor.LoadDeployConfigFromYaml(applyFlag.File)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		ctx, cancel := context.WithTimeout(context.Background(), supervisor.DefaultTimeout)
		defer cancel()
		client, err := getClient(ctx, applyFlag.Host)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		resp, err := client.Apply(ctx, &supervisor.ApplyReq{Config: &dc})
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		logger.Info(resp.Msg)
	},
}

func init() {
	applyCmd.PersistentFlags().StringVar(&applyFlag.Host, "host", "127.0.0.1:50060", "host")
	applyCmd.PersistentFlags().StringVar(&applyFlag.File, "file", "", "deploy yaml")
	rootCmd.AddCommand(applyCmd)
}