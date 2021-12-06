package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func cron_demo() {
	spec := "0 17/1 18 * * *" // 每天凌晨2：30
	c := cron.New()
	c.AddFunc(spec, callYourFunc)
	c.Start()
	select {}
}

func callYourFunc() {
	fmt.Println("callYourFunc come here.")
}

// 每天凌晨2：30就会调用一次callYourFunc函数
