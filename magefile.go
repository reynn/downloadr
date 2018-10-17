// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"io/ioutil"
	"os"
	"os/exec"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(BuildPlugins)
	fmt.Println("Building main...")
	cmd := exec.Command("go", "build", "-o", "bin/downloadr", "cmd/main.go")
	return cmd.Run()
}

func BuildPlugins() error {
	plugins, err := ioutil.ReadDir("plugins")
	if err != nil {
		return err
	}
	for _, p := range plugins {
		if !p.IsDir() {
			continue
		}
		fmt.Printf("Building plugin %s\n", p.Name())
		cmd := exec.Command(
			"go",
			"build",
			"--buildmode=plugin",
			"-o",
			fmt.Sprintf("plugins/%s.so", p.Name()),
			fmt.Sprintf("plugins/%s/main.go", p.Name()),
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		e := cmd.Run()
		if e != nil {
			fmt.Printf("Failed to build plugin %s, %v\n", p.Name(), e)
		}
	}
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("bin")
}
