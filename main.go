package main

import (
	"fmt"
	"time"

	"github.com/NorwayLobster/gomodone"
	"github.com/NorwayLobster/moduletest"
	"github.com/gin-gonic/gin"
	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
)

func generate(a uint32, b uint32) uint64 {
	fmt.Println("generate------start-----------------------------------------")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("a: %b\n", a)
	fmt.Printf("a uint64: %b\n", uint64(a))
	fmt.Printf("a<<32 uint64: %b\n", uint64(a)<<32)
	fmt.Printf("a: %x\n", a)
	fmt.Printf("a: %c\n", a)

	fmt.Printf("b: %d\n", b)
	fmt.Printf("b: %b\n", b)
	fmt.Printf("b uint64: %b\n", uint64(b))
	fmt.Printf("b: %x\n", b)
	fmt.Printf("b: %c\n", b)

	fmt.Printf("a<<32 uint64 +b : %b\n", (uint64(a)<<32)|uint64(b))
	fmt.Printf("a<<32 uint64 +b : %d\n", (uint64(a)<<32)|uint64(b))
	var score uint64 = (uint64(a) << 32) | uint64(b)
	fmt.Printf("score: %d\n", score)
	// for i := 0; i < 32; i++ {
	// 1 << i
	// score = score
	// }
	fmt.Println("generate------end-----------------------------------------")
	return score
}

func degenerate(score uint64) (uint32, uint32) {
	fmt.Println("degenerate------start-----------------------------------------")
	fmt.Printf("score: %d\n", score)
	var a uint32 = uint32(score >> 32)
	var b uint32 = uint32((score << 32) >> 32)
	fmt.Printf("a: %d\n", a)
	fmt.Printf("a: %b\n", a)
	// fmt.Printf("a uint64: %b\n", uint64(a))
	// fmt.Printf("a<<32 uint64: %b\n", uint64(a)<<32)
	fmt.Printf("a: %x\n", a)
	fmt.Printf("a: %c\n", a)

	fmt.Printf("b: %d\n", b)
	fmt.Printf("b: %b\n", b)
	fmt.Printf("b uint64: %b\n", uint64(b))
	fmt.Printf("b: %x\n", b)
	fmt.Printf("b: %c\n", b)

	// fmt.Printf("a<<32 uint64 +b : %b\n", (uint64(a)<<32)|uint64(b))
	// fmt.Printf("a<<32 uint64 +b : %d\n", (uint64(a)<<32)|uint64(b))
	// var score uint64
	fmt.Printf("score: %d\n", score)
	// for i := 0; i < 32; i++ {
	// 1 << i
	// score = score
	// }
	fmt.Println("degenerate----end-------------------------------------------")
	return a, b
	// return uint64(100)
}

func main() {
	moduletest.Hello()
	moduletest.Hello1()
	// moduletestV2.Proverb()
	// moduletest.Proverb()
	gomodone.SayHi("hello")
	fmt.Println(moduletest.Hello())
	r := gin.Default()
	r.GET("/ping", PongHandler)
	r.GET("/rankinglist/:name/:num", RankingListHandler)
	r.Run(":7080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	timestamp := time.Now().Unix()
	fmt.Printf("timestamp: %v\n", uint32(timestamp))
	var level uint32 = 55
	score := generate(uint32(99)-level, uint32(timestamp))
	fmt.Printf("after generate, score: %d\n", score)
	fmt.Println("-----------------------------------------------")
	a, b := degenerate(score)
	level = uint32(99) - a
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

}
