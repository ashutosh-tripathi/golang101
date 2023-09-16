package main

import "fmt"

func main() {
	fmt.Println("This is functions!!!")
	greeter()
	fmt.Println("sum of two is", addTwo(3, 2))
	a, b := reverse(2, 3)
	fmt.Println("reverse of 2,3 is", a, b)
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("reversed array is", reverseArray(arr))
	fmt.Println("The result of proAdd is", proAdd(1, 2, 3, 4, 5))
	x, y := namereturn(10, 20)
	fmt.Println("fetched values are", x, y)
}
func addTwo(a, b int) int {

	return a + b
}
func reverse(a, b int) (int, int) {
	return b, a
}
func greeter() {
	fmt.Println("hello from greeter!")
	func() {
		fmt.Println("This is self executing functions")
	}()
}
func reverseArray(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i <= len(a)/2; i++ {
		c, d := reverse(a[i], a[len(a)-i-1])
		b[i] = c
		b[len(a)-i-1] = d
	}
	return b
}
func proAdd(values ...int) int {

	sum := 0
	for _, a := range values {
		sum += a
	}
	return sum

}

// named return function
func namereturn(x, y int) (min, max int) {

	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}
