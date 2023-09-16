package main

import "fmt"

func num(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := num(nums...)
	c = square(c)
	for result := range c {
		fmt.Println("reslt", result)
	}
}
