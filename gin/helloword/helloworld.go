package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	helloWorld()
}

func helloWorld() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	engine.Run(":9000")
}
