package main

import (
	"fmt"
)

func generate1(a uint32, b uint32) uint64 {
	var score uint64 = (uint64(a) << 32) | uint64(b)
	return score
}

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

func degenerate1(score uint64) (uint32, uint32) {
	var a uint32 = uint32(score >> 32)
	var b uint32 = uint32((score << 32) >> 32)
	return a, b
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

func test() {

	timestamp := 1637050964
	// timestamp := time.Now().Unix()
	// var score uint64 = 216385415764
	fmt.Printf("timestamp: %v\n", uint32(timestamp))
	var level uint32 = 49
	score := generate(uint32(99)-level, uint32(timestamp))
	fmt.Printf("after generate, score: %d\n", score)
	fmt.Println("-----------------------------------------------")
	a, b := degenerate(score)
	level = uint32(99) - a
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("level: %d\n", level)
}
