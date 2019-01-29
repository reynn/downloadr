package main

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
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
		Use:   "plugins",
		Short: "Plugin management",
		Run: func(c *cobra.Command, aa []string) {
			fmt.Printf("There are %d plugins available\n", mana.GetCount())
		},
	}

	pluginsCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "list all plugins",
		Run: func(c *cobra.Command, aa []string) {
			displayPlugins(mana.GetScraperPlugins(), os.Stdout)
		},
	})

	return pluginsCmd
}

func displayPlugins(plugins []downloadr.ScraperPlugin, w io.Writer) {
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Version"})

	for _, p := range plugins {
		table.Append([]string{p.Name(), p.Version()})
	}

	table.Render()
}

func versionPlugin() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display version information",
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
