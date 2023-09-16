package main

import "fmt"

func main() {

	fmt.Println("this is struct")
	ashu := User{"Ashu", 28, "ashu@go.dev", true}
	fmt.Println("printing struct", ashu)
	fmt.Printf("printing %v", ashu)
	fmt.Printf("printing detail %+v", ashu)
	fmt.Println("printing age", ashu.Age)

}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Status bool   `json:"status"`
}
