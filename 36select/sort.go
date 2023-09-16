package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	m.Lock()
	a := []int{7, 2, 8, 4, 5}
	m.Unlock()
	sort.Ints(a)
	fmt.Println("a", a)

	go func() {
		for i, _ := range a {
			a[i] = a[i] + 1
		}
	}()
	go func() {
		for _, val := range a {
			fmt.Println("read value", val)
		}
	}()
	time.Sleep(3 * time.Second)
}
