package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"linkweek-go/db/repo"
	"linkweek-go/fetch"
)

var (
	amount int
)

func init() {
	fetchCmd.Flags().IntVarP(&amount, "amount", "a", 5, "Amount of top stories to fetch")
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch top stories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Fetching [%d] top stories ... ", amount)
		items := fetch.Fetch(amount)
		fmt.Println("ok.")
		fmt.Print("Persisting ... ")
		repo.Persist(items)
		fmt.Println("ok.")
	},
}
