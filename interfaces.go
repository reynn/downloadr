package downloadr

import (
	"github.com/spf13/cobra"
)

// ScraperPlugin will handle scraping a site to discover URLs to download
type ScraperPlugin interface {
	Register(rootCmd *cobra.Command, logger Logger) error
	Name() string
	Version() string
}

// PluginManager manage the load and usage of plugins
type PluginManager interface {
	Gather(rootCmd *cobra.Command) error
	GetScraperPlugins() []ScraperPlugin
	GetCount() int
}

type Logger interface {
	Info(f string, ff ...interface{})
	Debug(f string, ff ...interface{})
	Error(f string, ff ...interface{})
	Fatal(f string, ff ...interface{})
}
