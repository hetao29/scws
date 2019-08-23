package main

import (
	"github.com/gin-gonic/gin"
	"scws"
)

func init() {
}

func main() {
	scws_api := scws.New()
	scws_api.Set()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/words", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		words := c.DefaultQuery("key", "");
		scws_api.Get(words)
	})
	r.Run(":8020") // listen and serve on 0.0.0.0:8080
}
