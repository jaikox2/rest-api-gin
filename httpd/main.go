package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/jaikox2/rest-api-gin/httpd/handler/ping"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet("responese a string"))

	r.Run(":8080")
}
