package main

import (
	"github.com/reynn/downloadr/log"
	"github.com/reynn/downloadr/plugin"
	"github.com/spf13/cobra"
)

const (
	version = "0.1.0"
)

func main() {
	logger := log.New()
	pManager := plugin.New("plugins", logger)

	rootCmd := &cobra.Command{
		Use:   "downloadr",
		Short: "DownloadR is a plugable way to download from various places.",
		Long:  "DownloadR provides a simple and customizable way to get media from various places.",
		Run: func(c *cobra.Command, aa []string) {
			logger.Info("Welcome to the DownloadR")
		},
	}

	pManager.Gather(rootCmd)

	rootCmd.Execute()
}
