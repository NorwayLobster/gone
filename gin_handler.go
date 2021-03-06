/*
 * @Date: 2022-02-21 10:21:15
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-04-07 11:16:27
 * @FilePath: /gone/gin_handler.go
 */
package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

func PongHandler(ctx *gin.Context) {
	// ctx.QueryArray()
	name := ctx.Query("name")
	line := ctx.Query("line") // http://127.0.0.1:7080/ping?name=xiaocheng&line=10
	body, err := ctx.GetRawData()
	if err != nil {
	}
	fmt.Printf("body:%v\n", body)
	fmt.Printf("client ip:%v\n", ctx.ClientIP())
	// i1 := 10
	for i := 0; i < 10000; i++ {
		// for i := 0; i < math.MaxInt32; i++ {
		// logger.Debug("Trying to hit GET request for", zap.Int("val", i1))
		// logger.Debug("Trying to hit Pong Handler for", zap.String("line:", line), zap.Int("int val:", i1))
	}

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
