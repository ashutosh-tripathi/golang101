package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", 10)

	res := func1(ctx)
	fmt.Println("value of result is ", res)
	res = func2(ctx)
	fmt.Println("value of result is ", res)
	time.Sleep(3 * time.Second)

}
func func1(ctx context.Context) any {
	fmt.Println("in func 1 value is ", ctx.Value("user"))
	return ctx.Value("user")

}

func func2(ctx context.Context) any {
	fmt.Println("in func u user found is ", ctx.Value("user"))
	return "user"
}
