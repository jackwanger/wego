package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/repong/hardict"
)

// Auto versioning
var (
	env string
)

var port int

func init() {
	if env == "release" {
		runtime.GOMAXPROCS(runtime.NumCPU())
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	fmt.Println("Listening at", port)

	r := gin.Default()
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(fmt.Sprintf(":%d", port))
}

func validateEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	if hardict.ExistInvalidWord(text) {
		c.JSON(200, gin.H{"result": "false"})
	} else {
		c.JSON(200, gin.H{"result": "true"})
	}
}

func filterEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	text = hardict.ReplaceInvalidWords(text)
	c.JSON(200, gin.H{"result": text})
}
