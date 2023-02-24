package main

import (
	"github.com/juxuny/supervisor/log"
	"github.com/spf13/cobra"
	"os"
)

var logger = log.NewPrefix("[su]")

var (
	rootCmd = &cobra.Command{
		Use:   "supervisor-ctl",
		Short: "supervisor-ctl",
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}
