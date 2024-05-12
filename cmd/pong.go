package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pongCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping Cobra",
	Long:  "Ping Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pong")
	},
}
