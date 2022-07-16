package cmd

import (
	"linkweek-go/internal/fetch"
	"linkweek-go/internal/mapper"
	topstoryRepo "linkweek-go/internal/repos/topstory"
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
		log.Printf("Save...\n")
		topStories := mapper.ItemToDbModels(items)
		topstoryRepo.Save(topStories...)
		log.Printf("ok.\n")
	},
}
