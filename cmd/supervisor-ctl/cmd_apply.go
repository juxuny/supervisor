package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var applyFlag = struct {
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
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(baseFlag.Timeout)*time.Second)
		defer cancel()
		client, err := getClient(ctx, baseFlag.Host, baseFlag.CertFile)
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
	applyCmd.PersistentFlags().StringVar(&applyFlag.File, "file", "", "deploy yaml")
	rootCmd.AddCommand(applyCmd)
}
