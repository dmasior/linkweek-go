package fetch

import (
	"github.com/jarcoal/httpmock"
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	topStoriesResponseMock = `[ 28096019, 28090024, 28091750, 28094465 ]`
	firstItemResponseMock  = `{
"by": "abc",
"descendants": 46,
"id": 28096019,
"kids": [
28096722,
28096644
],
"score": 123,
"time": 1628314612,
"title": "story title",
"type": "story",
"url": "https://story.url"
}`
	secondItemResponseMock = `{
"by": "def",
"descendants": 77,
"id": 28090024,
"kids": [
28090025,
28096646
],
"score": 7,
"time": 1628314616,
"title": "2nd story title",
"type": "story nr 2",
"url": "https://story2.url"
}`
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
	items := Run(2)

	// assert
	assert.Equal(t, items[0].Id, 28096019)
	assert.Equal(t, items[1].Id, 28090024)
}
