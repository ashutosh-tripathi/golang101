//Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

package main

import "fmt"

func main() {
	var a = []int{100, 4, 200, 1, 3, 2}
	fmt.Println("lcu", lcu(a))
}

func lcu(nums []int) int {
	valmap := make(map[int]bool)

	for _, num := range nums {
		valmap[num] = true
	}
	var maxsum = 0
	var sum = 0
	for _, num := range nums {
		if valmap[num-1] != false {
			continue
		} else {
			sum = 1
			for valmap[num+1] != false {
				num = num + 1
				sum = sum + 1
			}
		}
		if sum > maxsum {
			maxsum = sum
		}
	}
	return maxsum
}
