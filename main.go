package main

import (
	"fmt"
	"strconv"

	"github.com/NorwayLobster/gomodone"
	"github.com/NorwayLobster/moduletest"
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

func module_demo() {
	moduletest.Hello()
	moduletest.Hello1()
	// moduletestV2.Proverb()
	// moduletest.Proverb()
	gomodone.SayHi("hello")
	fmt.Println(moduletest.Hello())
}

func main() {
	// parse_demo()
	// time_demo()
	// timer_demo()
	// ticker_demo()
	// cron_demo()
	// sort_demo()
	// receiver_demo()
	// interface_demo()
	// map_demo()
	// utf8_demo()
	// redis_demo()
	// test()
	// module_demo()
	// startHTTPServer()
	// uint32_demo()
	// app_demo()
	type_assertion_demo()
}
