package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("This is switch dice case")
	a := rand.Intn(6) + 1
	fmt.Println("You rolled", a)
	switch a {
	case 1, 6:
		fmt.Println("Please open")
	case 2, 3, 5:
		fmt.Println("Please move dice with ", a)
	default:
		fmt.Println("Please role again")

	}
}
