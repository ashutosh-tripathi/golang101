package main

import (
	"fmt"
	"net/url"
)

const weburl string = "http.lco.dev:3000/learn?coursename=temp&language=english"

// const weburl string = "abcdef*%&^*#)()"

func main() {
	fmt.Println("This is test for urls")
	fmt.Println("The url is:" + weburl)
	result, err := url.Parse(weburl)
	if err != nil {
		fmt.Println("Not a valid url:", err.Error())
	}
	queryResult := result.Query()
	fmt.Println("queryResult:", queryResult)
	for k, v := range queryResult {
		fmt.Println("Key and value are", k, v)
		fmt.Printf("Type of key and value are %T and %T", k, v)
	}

	fmt.Println("result is" + result.Scheme + " " + result.Host + " " + result.Path + " " + result.RawQuery + " " + result.Port())

}
