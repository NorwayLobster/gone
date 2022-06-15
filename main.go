/*
 * @Date: 2022-02-10 20:06:55
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-07-20 17:48:02
 * @FilePath: /gone/main.go
 */
package main

import (
	"fmt"
	"strconv"
	// "github.com/NorwayLobster/gomodone"
	// "github.com/NorwayLobster/moduletest"
	// cron "github.com/robfig/cron/v3"
	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
)

func parse_demo() {
	laterSignItemIdStr := "14216"
	laterSignItemId, ok := strconv.ParseUint(laterSignItemIdStr, 10, 32)
	// laterSignItemId, ok := strconv.ParseUint(laterSignItemIdStr, 0, 32)//the same
	if ok != nil {
		fmt.Printf("%v, %d\n", ok, laterSignItemId)
	}
	fmt.Printf("%d, %T\n", laterSignItemId, laterSignItemId)

	// a := "0b1111"
	// aBinary, ok := strconv.ParseUint(a, 0, 32)
	a := "1111"
	aBinary, ok := strconv.ParseUint(a, 8, 32)
	if ok != nil {
		fmt.Printf("%v, %d\n", ok, aBinary)
	}
	fmt.Printf("%v, %d, %b, %T\n", a, aBinary, aBinary, aBinary)
}

func main() {
	fmt.Printf("Hello world\n")
	fmt.Printf("Hello\n")
	// module_demo()
	// parse_demo()
	// time_demo()
	// timer_demo()
	// ticker_demo()
	// cron_demo()
	// sort_demo()
	// receiver_demo()
	// interface_demo()
	// map_demo()
	// sync_map_demo()
	// utf8_demo()
	// redis_demo()
	// test()
	// module_demo()

	//	// InitLogger(zapcore.DebugLevel)

	// InitLogger(zapcore.DebugLevel, 1, 10, 30, false)
	// defer logger.Sync()
	startHTTPServer()

	// uint32_demo()
	// app_demo()
	// type_assertion_demo()
	// null_interface_demo()
	// mysqlcli_demo()
	// rand.Seed(time.Now().Unix())
	// rand.Seed(1)
	// rand_demo()
	// atomic_demo()
	// sentry_demo()
	// zerolog_demo()
	// context_demo1()
	// a := Sum(1, 32)
	// fmt.Printf("sum:%d\n", a)
	// profiling_server_demo()
	// slice_delete_demo()
	// zap_demo()
	// InitLogger()
	// time_utc_gmt()
	// interface_embeded_into_struct()
	// prometheus_sdk()
	// runtime_stack()
	// fmt_format_demo()

	test_2()
}
