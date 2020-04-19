package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{
		"An Item", "An Item Value",
	})

	if len(feed.Items) == 0 {
		t.Errorf("Not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{
		"An Item", "An Item Value",
	})

	results := feed.GetAll()
	if len(results) == 0 {
		t.Errorf("Not got")
	}
}
