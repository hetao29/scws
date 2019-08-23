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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/words", func(c *gin.Context) {
		key := c.DefaultQuery("key", "");
		words:= scws_api.Get(key)
		c.JSON(200, gin.H{
			"message": "pong",
			"words": words,
		})
	})
	r.Run(":8020") // listen and serve on 0.0.0.0:8080
}
