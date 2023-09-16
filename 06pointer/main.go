package main

import "fmt"

func main() {
	fmt.Println("This is pointer")
	var num int = 23
	var ptr *int
	ptr = &num
	fmt.Println("The pointer address is", ptr)
	fmt.Println("The pointer value is ", *ptr)
	fmt.Println("The memory address of num is", &num)
	*ptr = *ptr * 2
	fmt.Println("updated value is ", *ptr)
	fmt.Println("update value num is ", num)
	more()
}
func more() {
	a := 1
	b := 2
	pa := &a
	pb := &b
	fmt.Println(" a and b are ", a, b)
	fmt.Println("pa and pb are", pa, pb)
	fmt.Println("value of pa and pb", *pa, *pb)

}
