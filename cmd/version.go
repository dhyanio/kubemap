package cmd

import (
	"fmt"
	"github.com/dhyanio/kubemap/constant"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kubemap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s\n", constant.AppName, constant.AppVersion, runtime.GOOS)
	},
}
