package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// go greeter("Hello")
	// greeter("World")
	// time.Sleep(3 * time.Millisecond)
	sites := []string{
		"http://www.example.com",
		"http://www.google.com",
		"http://www.github.com",
	}
	for _, site := range sites {
		go checkStatus(site)
		wg.Add(1)
	}
	wg.Wait()

}


//	func greeter(s string) {
//		for i := 0; i < len(s); i++ {
//			fmt.Println(s)
//		}
//	}
func checkStatus(s string) {
	defer wg.Done()
	res, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d status for call to %s \n", res.StatusCode, s)
}
