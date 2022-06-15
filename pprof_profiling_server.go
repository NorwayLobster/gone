/*
 * @Date: 2022-03-11 14:45:08
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-11 18:25:52
 * @FilePath: /gone/pprof_profiling_server.go
 */
package main

// 1. 引入 net/http/pprof 包
import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
// wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=5

// go tool  pprof --text http://10.24.14.37:58082/debug/pprof/heap
// curl --output heap.out.2 http://10.24.14.37:58082/debug/pprof/heap
// go tool pprof --http  localhost:9990 heap.out
// go tool pprof http://10.24.14.14:55500/debug/pprof/heap

// heap profile: 96(inused_objects): 1582948832(inused_bytes) [21847(allocated_objects): 15682528480(allocted_bytes)] @ heap/1048576

// 2. 开启 http 服务
func profiling_server_demo() {
	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe() error", err)
	}
}
