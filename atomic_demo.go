package main

import (
	"fmt"
	"sync/atomic"
)

var a int32 = 1
var b int32 = 100

func atomic_demo() {
	fmt.Printf("a:%d\n", a)
	atomic.AddInt32(&a, 1)
	fmt.Printf("a:%d\n", a)
	fmt.Printf("b:%d\n", b)
	fmt.Println()
	old_a := atomic.SwapInt32(&a, b)
	fmt.Printf("a:%d\n", a)
	fmt.Printf("old_a:%d\n", old_a)
	fmt.Printf("b:%d\n", b)

	atomic.Value()
}
