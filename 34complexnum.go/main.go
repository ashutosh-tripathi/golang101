package main

import "fmt"

func main() {
	c1 := 3 + 2i
	c2 := complex(5, 4)
	fmt.Printf("type of c1 %T", c1)
	fmt.Printf("type of c2 %T", c2)

	fmt.Println("printing complex", c1)
	fmt.Println("printing complex", c2)
	c3 := c1 + c2
	fmt.Println("sum ", c3)

	slice := make([]int, 2, 4)
	fmt.Println("slice is ", len(slice))

	callappend(slice)
	fmt.Println("slice is ", len(slice))

}
func callappend(a []int) {
	a = append(a, 10, 20, 30)

}
