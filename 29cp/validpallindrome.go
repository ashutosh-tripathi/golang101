// Given a string s, return true if it is a palindrome, or false otherwise.
// case1:"A man, a plan, a canal: Panama"
package main

import "fmt"

func main() {

	a := " "
	fmt.Println(isPalindrome(a))
}

func isPalindrome(s string) bool {
	var cleaned []rune

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			cleaned = append(cleaned, c)
		} else if c >= 'A' && c <= 'Z' {
			cleaned = append(cleaned, c+'a'-'A')
		}
	}
	fmt.Println("cleaned is ", string(cleaned))
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-i-1] {
			return false
		}
	}
	return true
}
