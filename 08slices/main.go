package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is slices!!")

	var a []string

	fmt.Printf("The type of array a is %T", a)
	fmt.Println("The length of a is", len(a))
	a = append(a, "a", "b")
	fmt.Println("a", a)
	fmt.Println("The length of a is", len(a))
	hs := make([]int, 4)
	hs[0] = 1
	hs[1] = 10
	hs[2] = 3
	hs[3] = 7
	hs = append(hs, 4, 8, 2)
	fmt.Println("hs", hs)
	sort.Ints(hs)
	fmt.Println("hs", hs)
	var index int = 2
	hs = append(hs[:index], hs[index+1:]...)
	fmt.Println("hs", hs)

}
