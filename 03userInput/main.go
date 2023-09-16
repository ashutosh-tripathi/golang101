package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is a test for user input")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("user entered", value)

	//using scanf to read user input
	var intval int
	fmt.Scanf("%d", &intval)
	fmt.Println("user entered input value", intval)
	fmt.Printf("type of input value %T", intval)

}
