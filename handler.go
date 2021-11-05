package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PongHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
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
