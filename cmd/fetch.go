package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"linkweek-go/fetch"
)

var (
	amount int
)

func init() {
	fetchCmd.Flags().IntVarP(&amount, "amount", "a", 10, "Amount of top stories to fetch")
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch top stories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fetch.Run(amount))
	},
}
