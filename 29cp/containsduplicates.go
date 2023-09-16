package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("contains", containsDuplicates(nums))
}
func containsDuplicates(nums []int) bool {

	val := make(map[int]int)
	for _, num := range nums {
		if val[num] != 0 {
			return true
		} else {
			val[num] = 1
		}
	}
	return false
}
