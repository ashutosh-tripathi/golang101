package main

import "fmt"

func main() {
	fmt.Println("This is a loop")
	days := make([]string, 0)
	otherdays := []string{"Tue", "Wed", "Thu", "Fri", "Sat"}
	days = append(days, "Sun", "Mon")
	days = append(days, otherdays...)
	fmt.Println("days", days)
	for i := 0; i < len(days); i++ {
		fmt.Println("day today is", days[i])
	}
	fmt.Println("range in for loop")
	for i, a := range days {
		fmt.Println("in range ", i, "day is", a)
	}
	fmt.Println("for as while loop")
	rogue := 0
	for rogue < 10 {
		if rogue == 5 {
			goto temp
		}
		if rogue == 5 {
			break
		}
		fmt.Println("The rogue value is ", rogue)
		rogue++
	}
temp:
	fmt.Println("This is temp go to")
}
