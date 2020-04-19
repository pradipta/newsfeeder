package main

import (
	"fmt"
	"newsfeeder/cmd/httpd/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Heelllo", "How you doing"})

	fmt.Println(feed)
}
