package importer

import (
	"fmt"
	legacyRepo "linkweek-go/internal/repos/legacy"
	"linkweek-go/internal/repos/topstory"
	"log"
	"math"
	"sync"
)

func ImportTopStories() {
	log.Println("Import top stories...")
	batchSize := 1000
	count := legacyRepo.CountItems()
	batches := int(math.Ceil(float64(count) / float64(batchSize)))

	maxGorutines := 5
	goroutineLoops := int(math.Ceil(float64(batches) / float64(maxGorutines)))

	fmt.Printf("count = %d, batchSize = %d, batches = %d, maxGoroutines = %d, goroutineLoops = %d\n", count, batchSize, batches, maxGorutines, goroutineLoops)
	for i := 0; i < goroutineLoops; i++ {
		fmt.Printf("Starting batch fetch [%d], \n", i)
		var wg sync.WaitGroup
		wg.Add(maxGorutines)
		for j := 0; j < maxGorutines; j++ {

			go func(i, j int) {
				innerOffset := j * batchSize
				offset := i*batchSize*maxGorutines + innerOffset
				if offset > count {
					// Looks like last batch run, too high offset. Nothing to fetch, exiting
					wg.Done()
					return
				}
				items := legacyRepo.GetItems(uint(offset), uint(batchSize))
				topstory.Save(items...)
				wg.Done()
			}(i, j)
		}

		wg.Wait()
	}
	log.Println("done.")
}
