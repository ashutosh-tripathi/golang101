//Problem statement: what number of words should I remove from string to make it a pallindrome

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abcb"
	for i := 0; i < len(s); i++ {

		if isPallindrome(s[i:]) {
			fmt.Println("number of chars to make it a pallindrome is", i)
			break
		}
	}
}

func isPallindrome(s string) bool {

	var isPallindrome bool = true



	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		fmt.Println("i and j are", s[i], s[j])
		if s[i] != s[j] {
			isPallindrome = false
		}
	}
	fmt.Println("String is pallindrome? ", isPallindrome)
	return isPallindrome
}
