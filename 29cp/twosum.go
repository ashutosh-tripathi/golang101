//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

package main

import "fmt"

func main() {
	var a = []int{3, 3}
	c := twosum(a, 6)
	fmt.Println("the result is", c)

}

func twosum(a []int, sum int) []int {
	valmap := make(map[int]int)
	for i, v := range a {
		if valmap[sum-v] != 0 {
			return []int{(valmap[sum-v] - 1), i}
		} else {
			valmap[v] = i + 1
		}
	}
	return []int{0, 0}
}
