package newsfeed_test

import (
	"testing"

	"github.com/jaikox2/rest-api-gin/platform/newsfeed"
)

func TestAdd(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"An Item", "With body"})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"An Item", "With body"})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
