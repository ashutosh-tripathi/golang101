package main

import "fmt"

func main() {

	defertest()

}

func defertest() int {

	fmt.Println("This is first line")

	defer fmt.Println("This is second line")
	defer fmt.Println("Will I come at last or second last?: second last it is")
	defer fmt.Println("checking where will I come")

	fmt.Println("This is third line")
	deferwithfor()

	return 0

	// defer fmt.Println("This is fourth line")
}
func deferwithfor() {

	for i := 0; i < 5; i++ {
		defer fmt.Println("The i is", i)
	}

}
