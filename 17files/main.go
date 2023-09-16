package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("These are files!!")

	file, err := os.Create("./temp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.WriteString(file, "Enter text here...")

	text, err := ioutil.ReadFile("./temp.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("The text read from the file..", string(text))
}
