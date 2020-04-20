package main

import (
	"newsfeeder/cmd/httpd/handler"
	"newsfeeder/cmd/httpd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	feed := newsfeed.New()

	r.GET("/ping", handler.PingGet())
	r.GET("/feed/get", newsfeed.NewsGet(feed))
	r.POST("/feed/post", newsfeed.AddNews(feed))

	r.Run()
}
