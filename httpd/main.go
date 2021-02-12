package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/jaikox2/rest-api-gin/httpd/handler/ping"
	"github.com/jaikox2/rest-api-gin/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet("Test pong"))
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run(":8080")
}
