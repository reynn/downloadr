package downloadr

import "github.com/spf13/cobra"

type DownloadPlugin interface {
	Register(rootCmd *cobra.Command, logger Logger) error
	GetName() string
	GetVersion() string
}

type PluginManager interface {
	Gather(rootCmd *cobra.Command) error
}

type Logger interface {
	Info(f string, ff ...interface{})
	Debug(f string, ff ...interface{})
	Error(f string, ff ...interface{})
	Fatal(f string, ff ...interface{})
}
