package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/moehandi/campaign/api/runner"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve server application",
	Long: `Serve application for run server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init server run ...")
		runner.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
