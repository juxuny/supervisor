package main

import (
	"context"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var applyFlag = struct {
	FileList []string
}{}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply",
	Run: func(cmd *cobra.Command, args []string) {
		if len(applyFlag.FileList) == 0 {
			logger.Error("missing file")
			os.Exit(-1)
		}

		for _, h := range baseFlag.Host {
			func() {
				for _, deployFile := range applyFlag.FileList {
					func() {
						dc, err := supervisor.LoadDeployConfigFromYaml(deployFile)
						if err != nil {
							logger.Error(err)
							os.Exit(-1)
						}
						ctx, cancel := context.WithTimeout(context.Background(), time.Duration(baseFlag.Timeout)*time.Second)
						defer cancel()
						client, err := getClient(ctx, h, baseFlag.CertFile)
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
					}()
				}
			}()
		}
	},
}

func init() {
	initBaseFlag(applyCmd)
	applyCmd.PersistentFlags().StringSliceVar(&applyFlag.FileList, "file", []string{}, "deploy yaml")
	rootCmd.AddCommand(applyCmd)
}
