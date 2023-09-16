package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	urls := []string{
		"http://127.0.0.1:8081/getAllUser",
		"http://127.0.0.1:8081/getUser/ashu",
		"http://127.0.0.1:8081/temp",
	}
	results := make(chan string)
	for _, url := range urls {
		go fetchRequest(ctx, url, results)
	}
	for range urls {
		fmt.Println(<-results)
	}

}

func fetchRequest(ctx context.Context, url string, results chan<- string) {
	if strings.Contains(url, "temp") {
		time.Sleep(5 * time.Second)
		fmt.Println("I come after 5 seconds")
	}
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error getting response from %s: %v", url, err)
	}
	defer resp.Body.Close()
	results <- fmt.Sprintf("Response from %s: %v", url, resp.StatusCode)
}
