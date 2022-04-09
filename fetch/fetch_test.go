package fetch

import (
	_ "embed"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//go:embed test/mock/http_response/top_story_response_1.json
	topStoriesResponseMock string
	//go:embed test/mock/http_response/item_response_1.json
	firstItemResponseMock string
	//go:embed test/mock/http_response/item_response_2.json
	secondItemResponseMock string
)

func TestFetch(t *testing.T) {
	// arrange
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://hacker-news.firebaseio.com/v0/topstories.json",
		httpmock.NewStringResponder(200, topStoriesResponseMock))

	httpmock.RegisterResponder("GET", "https://hacker-news.firebaseio.com/v0/item/28096019.json",
		httpmock.NewStringResponder(200, firstItemResponseMock))

	httpmock.RegisterResponder("GET", "https://hacker-news.firebaseio.com/v0/item/28090024.json",
		httpmock.NewStringResponder(200, secondItemResponseMock))

	// act
	items := Fetch(2)

	// assert
	assert.Equal(t, items[0].Id, uint(28096019))
	assert.Equal(t, items[1].Id, uint(28090024))
	assert.Equal(t, 2, len(items))
}
