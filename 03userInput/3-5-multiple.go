//programming problem that gives the sum of multiples of 3 and 5 below a number N
// T - number of iterations
// N - number per iternation
// for ex: 2
//10
//100
//output:
//23
//2318

// solution
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t, n int
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	// fmt.Println("fethced val", val)
	t, _ = strconv.Atoi(val)
	// fmt.Println("fetched t", t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		fmt.Println(getSum(n))
	}
}
func getSum(n int) int {
	var sum int = 0
	for i := 3; i < n; i += 3 {
		// fmt.Println("adding ", i)
		sum += i
	}
	fmt.Println("in 5 loop")
	for i := 5; i < n; i += 5 {
		// fmt.Println("adding ", i)
		if i%3 != 0 {
			sum += i
		}
	}

	return sum

}
