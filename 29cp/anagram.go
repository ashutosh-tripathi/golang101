// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
package main

import "fmt"

func main() {
	fmt.Println("is valid anagram?", validAnagram("anagram", "nagram"))

}
func validAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	smap := make(map[rune]int)
	tmap := make(map[rune]int)
	for i, _ := range s {
		smap[rune(s[i])] = smap[rune(s[i])] + 1
		tmap[rune(t[i])] = tmap[rune(t[i])] + 1
	}

	if len(smap) >= len(tmap) {
		for k, _ := range smap {
			if smap[k] != tmap[k] {
				return false

			}
		}
	} else {
		for k, _ := range tmap {
			if smap[k] != tmap[k] {
				return false
			}
		}

	}
	return true
}
