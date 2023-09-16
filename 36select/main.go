package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan int, 5)
	chan2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 5
	}()
	go func() {
		time.Sleep(time.Second * 1)
		chan2 <- 5
	}()

	select {
	case <-chan1:
		fmt.Println("hello from channel1")
	case <-chan2:
		fmt.Println("Hello from channel2")
	case <-time.After(4 * time.Second):
		fmt.Println("timed out")

	}

}
