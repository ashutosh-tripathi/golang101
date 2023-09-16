package main

import "fmt"

type User struct {
	name   string
	age    int
	email  string
	number string
}

func (u *User) getAge() int { return u.age }

func main() {

	user := User{"ashu", 28, "ashu@go.dev", "93251"}
	fmt.Println("age od user is", user.getAge())
	fmt.Println("user name is", user.name)
}
