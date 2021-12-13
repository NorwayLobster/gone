package main

import (
	"fmt"
	"log"
	"time"

	// "github.com/robfig/cron/v3"
	"github.com/robfig/cron"
)

func cron_demo() {
	// demo3()
	// demo1()
	// demo0()
	demo4()
	// demo2()
}

func demo4() {
	spec := "2 30 18 7 12 *" // 12月7日18:30:02执行
	c := cron.New()
	c.AddFunc(spec, callYourFunc)
	c.Start()
	select {}
}

func demo() {
	spec := "0 17/1 18 * * *" // 每天凌晨2：30
	c := cron.New()
	c.AddFunc(spec, callYourFunc)
	c.Start()
	// select {}
}

func callYourFunc() {
	// fmt.Println("callYourFunc come here.")
	log.Println("callYourFunc come here.")
}

// 每天凌晨2：30就会调用一次callYourFunc函数

func demo0() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("0 0", func() {
		// c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}

func demo1() {
	log.Println("Starting...")
	i := 0
	c := cron.New()
	spec := "*/1 * * * * *"
	// spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("start", i)
	})
	c.Start()
	select {} //阻塞主线程不退出
}

func text() {
	log.Println("text")
}

func demo2() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { text() })
	c.Start()
	select {}
}

func demo3() {
	cron2 := cron.New() //创建一个cron实例

	//执行定时任务（每1秒执行一次）
	err := cron2.AddFunc("*/1 * * * * *", print5)
	if err != nil {
		fmt.Println(err)
	}

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

//执行函数
func print5() {
	fmt.Println("每1s执行一次cron")
}
