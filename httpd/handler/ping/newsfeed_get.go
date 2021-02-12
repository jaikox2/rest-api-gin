package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaikox2/rest-api-gin/platform/newsfeed"
)

func NewsFeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, feed.GetAll())
	}
}
