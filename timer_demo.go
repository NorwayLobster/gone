package main

import (
	"fmt"
	"time"
)

// timer.After()
// NewTimer()
// timer.AfterFunc()

func timer_demo1() {
	f := func() { fmt.Println("I Love You!") }
	time.AfterFunc(time.Second*2, f)

	time.Sleep(time.Second * 4)
}

func timer_demo2() {
	f := func() { fmt.Println("I Love You!") }
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-timer.C:
			go f()
			timer.Reset(time.Second * 2)
		}
	}
}

func timer_demo3() {
	f := func() { fmt.Println("I Love You!") }
	for {
		select {
		case <-time.After(2 * time.Second):
			go f()
		}
	}
}

func timer_demo() {
	f := func() { fmt.Println("I Love You!") }
	timerForStop := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-time.After(2 * time.Second):
			go f()
		case <-timerForStop.C:
			return //stop
		}
	}
}

// NewTickerr()
func ticker_demo1() {
	f := func() { fmt.Println("I Love You!") }
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		go f()
	}
	// ticker.Stop()
}

func ticker_demo2() {
	f := func() { fmt.Println("I Love You!") }
	for range time.Tick(1 * time.Second) {
		go f()
	}
}

func ticker_demo() {
	f := func() { fmt.Println("I Love You!") }
	for {
		select {
		// case ticker <- time.Tick(1 * time.Second):
		case <-time.Tick(1 * time.Second):
			go f()
		}
	}
}
