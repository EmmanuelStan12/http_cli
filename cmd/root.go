package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "httpace-cli",
	Short: "httpace-cli is a tool for making HTTP requests",
	Long:  `A Fast and Flexible HTTP CLI tool built with love in Go.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
