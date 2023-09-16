package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is demo for if else!!")

	// rand.Seed(time.Now().UnixNano())

	a := rand.Intn(200)
	fmt.Println("number is ", a)
	if a < 100 {
		fmt.Println("number is less than 100")
	} else if a > 100 {
		fmt.Println("number is greater than 100")
	} else {
		fmt.Println("number is 100")
	}

}
