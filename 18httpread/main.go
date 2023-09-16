package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("This is http read")
	response, err := http.Get("https://lco.dev")
	defer response.Body.Close()
	if err != nil {
		fmt.Println("got error: ", err)
	}
	fmt.Println("status code: ", response.Status)
	fmt.Println("response header", response.Header)
	databytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("response body", string(databytes))
}
