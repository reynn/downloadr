package main

import (
	"github.com/reynn/downloadr"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

// ExamplePlugin adheres to the DownloadPlugin interface within the downloadr package.
type ExamplePlugin struct{}

// Register registers the plugin with Cobra as a subcommand, can take advantage of anything within cobra to make robust commands
func (v *ExamplePlugin) Register(rootCmd *cobra.Command, logger downloadr.Logger) error {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "example",
		Short: "An example of how to create a plugin",
		Long:  "An example of plugins for Downloadr, this does nothing on it's own",
		Run: func(c *cobra.Command, aa []string) {
			logger.Info("ExamplePlugin with arguments %v", aa)
		},
	})
	return nil
}

// GetName should return the name of the plugin
func (v *ExamplePlugin) Name() string {
	return "ExamplePlugin"
}

// GetVersion should return the version of the plugin
func (v *ExamplePlugin) Version() string {
	return version
}

var Plugin ExamplePlugin // This must be in the plugin like this so that we can load the plugin
