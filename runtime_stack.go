/*
 * @Date: 2022-03-09 16:46:50
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-09 19:19:56
 * @FilePath: /gone/runtime_stack.go
 */

package main

import (
	"log"
	"runtime"
)

func runtime_stack() {

	runtime.GC()
	// runtime.Caller()

	//print mem status
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)

	//print stack
	buf := make([]byte, 1<<20)
	runtime.Stack(buf, true)
	log.Printf("\n%s", buf)

	// var sr runtime.StackRecord
	// runtime.MemProfileRecord()

	//stack内存管理
	// runtime.morestack
	// runtime.newstack()
	// runtime.stackalloc()
	// runtime.copystack()

	// runtime.shrinkfree()
	// runtime.stackfree()
}
