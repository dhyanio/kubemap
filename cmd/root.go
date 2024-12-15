package cmd

import (
	"fmt"
	"github.com/dhyanio/kubemap/constant"
	"os"

	"github.com/spf13/cobra"
)

const (
	longAppDesc = "Kubemap is a tfstate file visualizer"
	shortAppDesc = "Kubemap is a tfstate file visualizer"
	 
)

// RootCmd represents the base commnad
var RootCmd = &cobra.Command{
	Use: "kubemap",
	Short: shortAppDesc,
	Long: longAppDesc,

	Run: func(cmd *cobra.Command, args []string) {
		printInfo()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to executing command", err)
		os.Exit(1)
	}
}

// printInfo prints the application infromation
func printInfo() {
	fmt.Println("Name:", constant.AppName)
	fmt.Println("Version:", constant.AppVersion)
	fmt.Println("Author:", constant.AuthorName)
}