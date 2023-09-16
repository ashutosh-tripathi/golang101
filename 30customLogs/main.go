package main

import (
	"log"
	"os"
)

func main() {

	file, _ := os.Create("./app.log")
	log.SetOutput(file)
	log.Println("This is log file")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("This is log file2")
	log.

}
