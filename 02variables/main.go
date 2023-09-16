package main

import "fmt"

const LoginToken string = "login"

func main() {
	var username string = "username"
	fmt.Println("username: ", username)
	fmt.Printf("username is of type %T\n", username)

	var isLogged bool = false
	fmt.Println("is logged: ", isLogged)
	fmt.Printf("is logged is of type %T\n", isLogged)

	var value uint8 = 255
	fmt.Println("value", value)
	fmt.Printf("value is of type %T\n", value)

	var flvalue string
	fmt.Println("fl value", len(flvalue))
	fmt.Printf("flvalue  is of type %T \n", flvalue)

	s := 255
	fmt.Printf("s is of type %T \n", s)

	fmt.Println("login token", LoginToken)

}
