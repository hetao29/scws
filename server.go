package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
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
		key := c.DefaultQuery("key", "")
		words := scws_api.Get(key)
		c.JSON(200, gin.H{
			"message": "pong",
			"words":   words,
		})
	})
	r.GET("/reload", func(c *gin.Context) {
		scws_api.Reload()
		c.JSON(200, gin.H{
			"message": "pong",
			"reload":  true,
		})
	})
	var bindaddr *string
	bindaddr = flag.String("b", "0.0.0.0:80", "listen port")
	flag.Parse()
	log.Println("bind addr:%v", *bindaddr)
	err := r.Run(*bindaddr) // listen and serve on 0.0.0.0:8080
	log.Println("err:%v", err)
}
