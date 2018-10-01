package main

import (
	"fmt"
	"github.com/reynn/downloadr"
	"github.com/reynn/downloadr/log"
	"github.com/reynn/downloadr/plugin"
	"github.com/spf13/cobra"
)

const (
	version = "0.1.0"
)

func pluginCommand(mana downloadr.PluginManager) *cobra.Command {
	pluginsCmd := &cobra.Command{
		Use:"plugins",
		Run: func(c *cobra.Command, aa []string) {
			fmt.Printf("There are %d plugins available\n", mana.GetCount())
		},
	}

	pluginsCmd.AddCommand(&cobra.Command{
		Use:"list",
		Run: func(c *cobra.Command, aa []string) {
			for _, p := range mana.GetPlugins(){
				fmt.Printf("Name: %s\tVersion: %s\n", p.GetName(), p.GetVersion())
			}
		},
	})

	return pluginsCmd
}

func versionPlugin() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(c *cobra.Command, aa []string) {
			fmt.Printf("DownloadR version: %s\n", version)
		},
	}
}

func main() {
	logger := log.New()
	pManager := plugin.New("plugins", logger)

	rootCmd := &cobra.Command{
		Use:   "downloadr",
		Short: "DownloadR is a plugable way to download from various places.",
		Long:  "DownloadR provides a simple and customizable way to get media from various places.",
	}

	pManager.Gather(rootCmd)

	rootCmd.AddCommand(versionPlugin(), pluginCommand(pManager))

	rootCmd.Execute()
}
