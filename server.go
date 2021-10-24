package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"path"
	"scws"
	//"scws/config"
)

func init() {
}

func main() {
	_file, _ := exec.LookPath(os.Args[0])
	_pwd, _ := path.Split(_file)
	os.Chdir(_pwd)

	//c := flag.String("c", "config.json", "Specify the configuration file.")
	bindaddr := flag.String("b", "0.0.0.0:80", "listen port")
	flag.Parse()
	file, err := os.Open("../etc/config.json")
	if err != nil {
		log.Fatal("error: can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := scws.Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("error: can't decode config JSON: ", err)
	}

	scws_api := scws.New(Config)
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
	log.Println("notice: bind addr:%v", *bindaddr)
	err = r.Run(*bindaddr) // listen and serve on 0.0.0.0:8080
	log.Println("error: %v", err)
}
