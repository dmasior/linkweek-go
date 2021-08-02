package cmd

import (
	"github.com/spf13/cobra"
	"linkweek-go/fetch"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch top stories",
	Run: func(cmd *cobra.Command, args []string) {
		// wip: get count from fetch cmd args
		fetch.Run(10)
	},
}
