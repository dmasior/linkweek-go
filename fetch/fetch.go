package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type topStoryId int

type item struct {
	Id          int    `json:"id"`
	Deleted     int    `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Kids        []int  `json:"kids"`
	Url         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Descendants int    `json:"descendants"`
}

func Run(count int) {
	ids := fetchTopIds(count)

	for _, value := range fetchItems(ids) {
		fmt.Println(value)
	}
}

func fetchItems(ids []topStoryId) []item {
	var items []item

	reqCh := make(chan item)
	for _, itemId := range ids {
		go fetchItem(itemId, reqCh)
		items = append(items, <-reqCh)
	}

	return items
}

func fetchItem(id topStoryId, reqCh chan item) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	client := http.Client{
		Timeout: time.Second * 3,
	}

	resp, getErr := client.Get(url)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	var item item

	jsonErr := json.Unmarshal(body, &item)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if item.Url == "" {
		setItemUrl(&item)
	}

	reqCh <- item
}

func setItemUrl(i *item) {
	i.Url = fmt.Sprintf("https://news.ycombinator.com/item?id=%d", i.Id)
}

func fetchTopIds(amount int) []topStoryId {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	client := http.Client{
		Timeout: time.Second * 2,
	}

	resp, getErr := client.Get(url)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	var topStoryIds []topStoryId

	jsonErr := json.Unmarshal(body, &topStoryIds)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// return only first {amount} elements
	return topStoryIds[:amount]
}
