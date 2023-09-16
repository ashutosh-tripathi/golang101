package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url string = "http://localhost:8080"

func main() {

	fmt.Println("This is handling get Request")
	handleGetRequest("/getUser?name=Ashu&age=28")
	// handlePostRequest("/createUser")
}

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Phonenumber string `json:"phonenumber"`
}

var user = User{Name: "Ashu", Email: "ashu@dev.com", Age: 28, Phonenumber: "12345"}

func handleGetRequest(queryPath string) {

	response, err := http.Get(url + queryPath)
	if err != nil {
		fmt.Println("Got error: ", err)
	}
	defer response.Body.Close()
	fmt.Println("response status: ", response.Status)
	fmt.Println("response headers: ", response.Header)

	databytes, err := io.ReadAll(response.Body)
	fmt.Println("read body", string(databytes))

}
func handlePostRequest(queryPath string) {

	userbytes, _ := json.Marshal(user)

	reader := bytes.NewReader(userbytes)

	response, err := http.Post(url+queryPath, "application/json", reader)

	if err != nil {
		fmt.Println("got error while posting request: ", err)
	}
	defer response.Body.Close()
	fmt.Println("response statuse", response)
	readbytes, _ := io.ReadAll(response.Body)
	fmt.Println("response body" + string(readbytes))
}
