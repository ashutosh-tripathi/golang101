package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time now is" + time.Now().String())
	fmt.Println("date today is" + time.DateOnly)
	fmt.Println("only time now is" + time.TimeOnly)
	fmt.Println("formatted time is", time.Now().Format("01-02-2006"))
	createdDate := time.Date(2022, time.December, 10, 12, 10, 10, 10, time.Local)
	fmt.Println("created date", createdDate)
	moretime()
}
func moretime() {
	fmt.Println("epoch time ", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println("day", t.Day(), t.Month(), t.Year())
	time.Sleep(3 * time.Second)
	d, err := time.Parse("02 Jan 2006 15:04:05", time.Now().String())
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("date is ", d)
}
