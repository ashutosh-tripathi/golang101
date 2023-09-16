package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go performTask(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 1)
	fmt.Println("everything ends now")

}
func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task completed successfully")
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("in select default")
		}
	}
}
