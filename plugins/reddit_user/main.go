package main

import (
	"github.com/reynn/downloadr"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var (
	clientId     string
	clientSecret string
)

// RedditUser adheres to the DownloadPlugin interface within the downloadr package.
type RedditUser struct{}

// Register registers the plugin with Cobra as a subcommand, can take advantage of anything within cobra to make robust commands
func (v *RedditUser) Register(rootCmd *cobra.Command, logger downloadr.Logger) error {
	rCmd := &cobra.Command{
		Use:     "reddit",
		Aliases: []string{"ru"},
		Short:   "Get posts from a user on Reddit",
		Long:    "An example of plugins for Downloadr, this does nothing on it's own",
		Run: func(c *cobra.Command, aa []string) {
			logger.Info("RedditUser with arguments %v", aa)
		},
	}
	rCmd.PersistentFlags().StringVar(&clientId, "client-id", "", "A Reddit app client id")
	rCmd.PersistentFlags().StringVar(&clientSecret, "client-secret", "", "A Reddit app client secret")
	rootCmd.AddCommand(rCmd)
	return nil
}

// GetName should return the name of the plugin
func (v *RedditUser) Name() string {
	return "RedditUser"
}

// GetVersion should return the version of the plugin
func (v *RedditUser) Version() string {
	return version
}

var Plugin RedditUser // This must be in the plugin like this so that we can load the plugin
