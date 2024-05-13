package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	sampleprivategomod "github.com/vucong2409/sample-private-go-mod"
)

var pongCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping Cobra",
	Long:  "Ping Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(sampleprivategomod.Ping())
	},
}
