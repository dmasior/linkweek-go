package cmd

import (
	"fmt"
	legacyRepo "linkweek-go/internal/repos/legacy"
	"linkweek-go/internal/repos/topstory"
	"log"
	"math"
	"sync"

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
		batchSize := 1000
		count := legacyRepo.CountItems()
		batches := int(math.Ceil(float64(count) / float64(batchSize)))

		maxGorutines := 10
		goroutineLoops := int(math.Ceil(float64(batches) / float64(maxGorutines)))

		fmt.Printf("count = %d, batchSize = %d, batches = %d, maxGoroutines = %d, goroutineLoops = %d\n", count, batchSize, batches, maxGorutines, goroutineLoops)
		for i := 0; i < goroutineLoops; i++ {
			fmt.Printf("Starting batch fetch [%d]\n", i)
			var wg sync.WaitGroup
			wg.Add(maxGorutines)
			for j := 0; j < maxGorutines; j++ {

				go func(i, j int) {
					innerOffset := j * batchSize
					offset := i*batchSize*maxGorutines + innerOffset
					if offset > count {
						fmt.Printf("Looks like last batch run [%d]. Too high offset [%d], nothing to fetch, exiting\n", i, offset)
						wg.Done()
						return
					}
					log.Printf("fetching [i = %d, j = %d, offset = %d] \n", i, j, offset)
					items := legacyRepo.GetItems(uint(offset), uint(batchSize))
					topstory.Save(items...)
					log.Printf("saving [i = %d, j = %d, offset = %d, items = %d] \n", i, j, offset, len(items))
					wg.Done()
				}(i, j)
			}

			wg.Wait()
		}
	},
}
