package cmd

import (
	"linkweek-go/db/repo"
	"linkweek-go/fetch"
	"log"

	"github.com/spf13/cobra"
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
		log.Printf("Fetching [%d] top stories ...\n", amount)
		items := fetch.Fetch(amount)
		log.Printf("ok.\n")
		log.Printf("Persisting ...\n")
		repo.Persist(items)
		log.Printf("ok.\n")
	},
}
