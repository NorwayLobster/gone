package main

import (
	"fmt"
	"math"
)

func uint32_demo() {
	var u1 uint32 = 10
	// var i1 int = 10
	if u1 == 10 {
		fmt.Println(99 - uint32(49))
		fmt.Printf("true:u1-10=%d\n", u1-10)
	} else {
		fmt.Println("")
	}
}

func uint64Andfloat64() {
	var a uint64 = 12423343535
	var b float64 = float64(a)
	var c uint64 = uint64(b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	a = math.MaxUint64
	b = float64(a)
	c = uint64(b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
