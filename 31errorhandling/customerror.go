package main

import "fmt"

type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CustomError{"divide by zero"}
	} else {
		return a / b, nil
	}
}

func main() {
	val, err := divide(4, 4)
	if err != nil {
		fmt.Println("got error: ", err)
	} else {
		fmt.Println("val is", val)
	}
	val, err = divide(4, 0)
	if err != nil {
		fmt.Println("got error: ", err)
	} else {
		fmt.Println("val is", val)
	}
}
