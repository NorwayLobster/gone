package main

import (
	"fmt"

	"github.com/NorwayLobster/gomodone"
	// "github.com/NorwayLobster/moduletest"

	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
	"github.com/NorwayLobster/moduletest/v3/greetings"
	"github.com/NorwayLobster/moduletest/v3/hello"
	"github.com/winlinvip/mod_major_releases"
	v2 "github.com/winlinvip/mod_major_releases/v2"
)

//you should read // https://go.dev/learn/
// For a high-level module development workflow – which includes publishing – see Module release and versioning workflow.
//0.1 https://go.dev/doc/modules/release-workflow
// For an overview of module development, see Developing and publishing modules
//0.2 https://go.dev/doc/modules/developing

//1. https://go.dev/doc/tutorial/create-module
//2. https://go.dev/doc/modules/managing-source
//3. https://go.dev/doc/modules/managing-dependencies#naming_module
//4. https://go.dev/blog/module-compatibility
func module_demo() {
	module_demo1()
	mod_major_releases_demo()
}

func mod_major_releases_demo() {
	fmt.Println("Hello",
		mod_major_releases.Version(),
		v2.Version2(),
	)
}

func module_demo1() {
	greetings.Hello("greet to bili")
	hello.Hello1()
	// moduletestV2.Proverb()
	// moduletest.Proverb()
	gomodone.SayHi("hello")
	// fmt.Println(moduletest.Hello())

}
