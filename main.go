package main

import (
	"fmt"

	"github.com/NorwayLobster/gomodone"
	"github.com/NorwayLobster/moduletest"
	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
)

func module_demo() {
	moduletest.Hello()
	moduletest.Hello1()
	// moduletestV2.Proverb()
	// moduletest.Proverb()
	gomodone.SayHi("hello")
	fmt.Println(moduletest.Hello())
}

func main() {
	// time_demo()
	map_demo()
	// redis_demo()
	// test()
	// module_demo()
	// startHTTPServer()
}
