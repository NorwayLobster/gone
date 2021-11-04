package main

import (
	"github.com/NorwayLobster/gomodone"
	"github.com/NorwayLobster/moduletest"

	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	moduletest.Hello()
	moduletest.Hello1()
	// moduletestV2.Proverb()
	// moduletest.Proverb()
	gomodone.SayHi("hello")
	// fmt.Println(moduletest.Hello())
	r := gin.Default()
	r.GET("/ping", PongHandler)
	r.GET("/rankinglist/:name/:num", RankingListHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
