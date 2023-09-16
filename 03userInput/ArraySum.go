// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	n, _ := strconv.Atoi(val)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		val, _ = reader.ReadString('\n')
		val = strings.TrimSpace(val)
		arri, _ := strconv.Atoi(val)
		a[i] = arri
	}
	fmt.Println("array", a)

	val, _ = reader.ReadString('\n')
	val = strings.TrimSpace(val)
	reqSum, _ := strconv.Atoi(val)
	requiredSum(a, reqSum)
}

func requiredSum(a []int, reqSum int) {

	mapnum := make(map[int]int)
	for i := 0; i < len(a); i++ {
		fmt.Println("mapnum", mapnum)
		fmt.Println("a[i]", a[i])
		fmt.Println("mapnum value", mapnum[reqSum-a[i]])
		if mapnum[reqSum-a[i]] != 0 {
			fmt.Println("fetched values", mapnum[reqSum-a[i]], i)
			fmt.Println("values at", reqSum-a[i], a[i])
			return
		}
		mapnum[a[i]] = i
	}

}
