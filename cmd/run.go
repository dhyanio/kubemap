package cmd

import (
	"github.com/dhyanio/kubemap/route"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run kubemap",
	Run: func(cmd *cobra.Command, args []string) {
		route.RunRoute()
	},
}
