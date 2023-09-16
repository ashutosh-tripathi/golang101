package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go performTask(ctx, cancel)
	select {
	case <-ctx.Done():
		fmt.Println("Context timeout")
	}

}
func performTask(ctx context.Context, cancel context.CancelFunc) {
	// time.Sleep(3 * time.Second)
	fmt.Println("Task completed successfully")
	cancel()
}
