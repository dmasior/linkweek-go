package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type topStoryId int

type Item struct {
	Id          uint   `json:"id"`
	Deleted     int    `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Kids        []int  `json:"kids"`
	Url         string `json:"url"`
	Score       uint   `json:"score"`
	Title       string `json:"title"`
	Descendants int    `json:"descendants"`
}

var (
	Client *http.Client
)

func init() {
	Client = &http.Client{
		Timeout: time.Second * 3,
	}
}

func Fetch(count int) []Item {
	ids := fetchTopIds(count)

	items := fetchItems(ids)

	return items
}

func fetchItems(ids []topStoryId) []Item {
	var wg sync.WaitGroup
	items := make([]Item, 0, len(ids))

	reqCh := make(chan Item, len(ids))
	wg.Add(len(ids))

	for _, itemId := range ids {
		go fetchItem(&wg, itemId, reqCh)
	}

	wg.Wait()
	close(reqCh)

	for item := range reqCh {
		items = append(items, item)
	}

	return items
}

func fetchItem(wg *sync.WaitGroup, id topStoryId, reqCh chan Item) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := Client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var item Item

	err = json.Unmarshal(body, &item)

	if err != nil {
		log.Fatal(err)
	}

	if item.Url == "" {
		setHnItemUrl(&item)
	}

	reqCh <- item
	wg.Done()
}

func setHnItemUrl(i *Item) {
	i.Url = fmt.Sprintf("https://news.ycombinator.com/Item?id=%d", i.Id)
}

func fetchTopIds(amount int) []topStoryId {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := Client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var topStoryIds []topStoryId

	err = json.Unmarshal(body, &topStoryIds)

	if err != nil {
		log.Fatal(err)
	}

	// return only first {amount} elements
	return topStoryIds[:amount]
}
