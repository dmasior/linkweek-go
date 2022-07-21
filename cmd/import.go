package cmd

import (
	"fmt"
	"linkweek-go/internal/importer"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import top stories from legacy storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Import start")

		importer.ImportPublications()
		importer.ImportTopStories()
		importer.LinkPublicationTopStories()

		fmt.Println("Import done")
	},
}
