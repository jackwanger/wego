package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/repong/wego/dict"
)

// Auto versioning
var (
	env string
)

var port int
var dictPath string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	flag.StringVar(&dictPath, "dict", "", "Directory path. Multiple directories use comma separated string like a,b,c")
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	dict.Load(dictPath)
	fmt.Println("Listening at", port)

	r := gin.Default()
	r.GET("/validate", validateEndPoint)
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(fmt.Sprintf(":%d", port))
}

func validateEndPoint(c *gin.Context) {
	text := c.Query("message")
	if dict.ExistInvalidWord(text) {
		c.JSON(200, gin.H{"result": "false"})
	} else {
		c.JSON(200, gin.H{"result": "true"})
	}
}

func filterEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	text = dict.ReplaceInvalidWords(text)
	c.JSON(200, gin.H{"result": text})
}
