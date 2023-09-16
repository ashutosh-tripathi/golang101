package main

import (
	"fmt"
)

func main() {
	fmt.Println("These are maps")

	maps := make(map[string]string)
	maps["js"] = "javascript"
	maps["py"] = "python"
	maps["sol"] = "solidity"
	a, b := maps["py"]
	fmt.Printf("type of a and b are %T and %T\n", a, b)
	fmt.Println("values of a and b", a, b)
	fmt.Println("maps", maps)
	z := "python"
	x, t := 1, 2
	fmt.Println("x,t", x, t)
	s1 := 23
	runez := []rune(z)
	fmt.Println("runez", runez)
	fmt.Println("deleting from maps")
	delete(maps, "py")
	fmt.Println("maps", maps)
	for key, val := range maps {
		fmt.Println("key and val are", key, val)
	}
	arraymap := make(map[string][]int)
	arraymap["js"] = []int{1, 2, 3, 4}

	arraymap["py"] = []int{2, 3, 4, 5}
	fmt.Println("printing arraymap", arraymap)

}
