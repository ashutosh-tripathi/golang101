package main

import (
	"fmt"
	"net/http"

	"github.com/ashutosh-tripathi/mongoapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("String go server...")
	http.ListenAndServe(":8080", r)
}
