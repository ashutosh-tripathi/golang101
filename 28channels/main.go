package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is channels!")
	mychan := make(chan int, 5)
	wg := &sync.WaitGroup{}
	// mychan <- 1
	// fmt.Println(<-mychan)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		mychan <- 1
		mychan <- 2
		fmt.Println("inserted in mychan")
		wg.Done()
	}(mychan, wg)

	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Printf("got from mychain %d\n", <-ch)
		wg.Done()
	}(wg, mychan)
	wg.Wait()
}
