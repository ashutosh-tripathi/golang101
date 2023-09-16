//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

package main

import "fmt"

func main() {
	var a = []int{2, 7, 11, 15}
	c := twosum(a, 9)
	fmt.Println("the result is", c)

}

func twosum(numbers []int, target int) []int {

	for l, r := 0, len(numbers)-1; l < r; {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}

	}
	return []int{0, 0}

}
