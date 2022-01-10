package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func context_demo1() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	go handle_func(ctx)
	for {
	}
}
func handle_func(ctx context.Context) {
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		default:
			for {
			}
		}
	}
}
