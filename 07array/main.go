package main

import "fmt"

func main() {
	fmt.Println("Processing array")

	var a [4]string
	a[0] = "a1"
	a[1] = "a2"
	fmt.Println("a", a)
	b := make([]int, 4)
	b[0] = 1
	b[1] = 2
	fmt.Println("b", b)

}
