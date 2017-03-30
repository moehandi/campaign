package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/fatih/color"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print current active version",
	Run: func(cmd *cobra.Command, args []string) {
		color.Set(color.FgRed, color.Bold)

		fmt.Println(" running on Campaign-v1.0.0")
		color.Unset()
	},
}
