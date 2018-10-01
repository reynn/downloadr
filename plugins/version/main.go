package main

import (
	"fmt"

	"github.com/reynn/downloadr"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

type Version struct{}

func (v *Version) Register(rootCmd *cobra.Command, logger downloadr.Logger) error {
	rootCmd.AddCommand(&cobra.Command{
		Use: "version",
		Run: func(c *cobra.Command, aa []string) {
			fmt.Printf("DownloadR version: %s", version)
		},
	})
	return nil
}

func (v *Version) GetName() string {
	return "Version"
}

func (v *Version) GetVersion() string {
	return version
}

var Plugin Version
