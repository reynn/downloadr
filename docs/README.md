# DownloadR

## Project Goals

Goal is to create a easy to use application to download files from across the internet.
To do this the project will take advantage of Golang's plugin support to allow independent updates.


## Basic plugin

A basic plugin would look something like this:

```go
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
		Use: "example",
		Run: func(c *cobra.Command, aa []string) {
			logger.Info("ExamplePlugin with arguments %v", aa)
		},
	})
	return nil
}

// GetName should return the name of the plugin
func (v *ExamplePlugin) GetName() string {
	return "ExamplePlugin"
}

// GetVersion should return the version of the plugin
func (v *ExamplePlugin) GetVersion() string {
	return version
}

var Plugin ExamplePlugin // This must be in the plugin like this so that we can load the plugin
```

Once this is in place you can build the plugin using the following command:

`go build -buildmode=plugin -ldflags "-X main.version=0.2.0" -o plugins/example.so plugins/example/main.go`