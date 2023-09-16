package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to PizzaHub")
	fmt.Println("Please rate our pizza between 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	inputval, _ := reader.ReadString('\n')
	inputval = strings.TrimSpace(inputval)
	inpuint, err := strconv.Atoi(inputval)

	if err != nil {
		fmt.Println("got error while conversion:", err)
	} else {
		fmt.Println("your rating is:", inpuint)
	}
}
