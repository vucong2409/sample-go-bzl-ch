package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vucong2409/sample-go-bzl-ch/router"
)

var port int

func init() {
	startServerCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port of the server")
}

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		router.StartServer(port)
	},
}
