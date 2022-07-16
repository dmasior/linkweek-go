package cmd

import (
	"linkweek-go/internal/dataimport"
	"linkweek-go/internal/repos/topstory"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import top stories from legacy storage",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Import start...\n")
		offset := 0
		limit := 500
		for {
			items := dataimport.GetItems(uint(offset), uint(limit))
			if len(items) == 0 {
				break
			}
			topstory.Save(items...)
			log.Printf("Saved [%d] items\n", len(items))
			offset += limit

		}
		log.Printf("done \n")
	},
}
