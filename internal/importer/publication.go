package importer

import (
	"fmt"
	legacyRepo "linkweek-go/internal/repos/legacy"
	"linkweek-go/internal/repos/publication"
	"linkweek-go/internal/repos/topstory"
	"log"
)

func ImportPublications() {
	log.Println("Import publications...")
	publications := legacyRepo.GetPublications()
	publication.Save(publications...)
	log.Println("done")
}

func LinkPublicationTopStories() {
	log.Println("Link Publication <-> TopStories")
	topStoryPublications := legacyRepo.GetTopStoryPublications()
	for _, topStoryPublication := range topStoryPublications {
		pub := publication.Find(topStoryPublication.PublicationId)
		ts := topstory.FindByIds(topStoryPublication.TopStoryIds)
		pub.TopStories = ts
		publication.Save(pub)
		fmt.Printf("Linked publication [%d] stories\n", pub.ID)
	}
}
