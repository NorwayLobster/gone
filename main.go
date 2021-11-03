package main

import (
	"github.com/NorwayLobster/gomodone"
	"github.com/NorwayLobster/moduletest"
	"github.com/gin-gonic/gin"
)

func main() {
	moduletest.Hello()
	moduletest.Hello1()
	gomodone.SayHi("hello")
	// fmt.Println(moduletest.Hello())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
