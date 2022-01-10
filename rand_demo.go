package main

import (
	"fmt"
	"math/rand"
)

func rand_demo() {
	var s = []uint32{1, 2, 31, 10}
	num := len(s)
	fmt.Printf("num: %d\n", num)
	for i := 0; i < 10; i++ {
		index := rand.Intn(num)
		fmt.Printf("index: %d\n", index)
		fmt.Println(s[index])
		fmt.Println()
	}
}
