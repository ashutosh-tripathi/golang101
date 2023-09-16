package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}
func main() {
	val, err := divide(1, 0)
	if err != nil {
		fmt.Println("encountered error", err)
		fmt.Println(fmt.Errorf("encountered error: %v", err))
		s := fmt.Sprintf("fetched value is %d", val)
		fmt.Println("s", s)
		if err.Error() == "division by zero" {
			fmt.Println("right error thrown")
		}
		// panic(err)
		a := []int{1, 2, 3, 4, 5}
		fmt.Println("len  and capacity s", len(a), cap(a))
		a = a[:2]

		fmt.Println("len  and capacity s", a, len(a), cap(a))
		a = append(a, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}...)
		fmt.Println("len  and capacity s", a, len(a), cap(a))

	}
	fmt.Println("val", val)
}
