package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PongHandler(ctx *gin.Context) {
	// ctx.QueryArray()
	name := ctx.Query("name")
	line := ctx.Query("line") // http://127.0.0.1:8080/ping?name=xiaocheng&line=10
	body, err := ctx.GetRawData()
	if err != nil {
	}
	fmt.Printf("body:%v\n", body)
	fmt.Printf("client ip:%v\n", ctx.ClientIP())

	// ctx.Status(200)
	// ctx.String(200, "hello world")
	ctx.JSON(200, gin.H{
		"message": "pong",
		"name":    name,
		"line":    line,
	})
}

func RankingListHandler(ctx *gin.Context) {
	rankinglistName := ctx.Param("name")
	fmt.Printf("rankinglistName: %s\n", rankinglistName)
	numLine, err := strconv.Atoi(ctx.Param("num"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("numLine: %d\n", numLine)
	ctx.JSON(200, gin.H{
		"message": "pong",
		// "rankinglist": marsh.Json(),
		// "rankinglistName": rankinglistName,
		// "numLine":         numLine,
	})
}
