package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "sample-go-bzl-ch",
	Short: "Short desc!",
	Long:  "Long long desc!",
}

func init() {
	rootCmd.AddCommand(pongCmd)
	rootCmd.AddCommand(startServerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
