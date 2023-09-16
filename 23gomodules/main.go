package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This iz go module")
	r := mux.NewRouter()
	r.HandleFunc("/home", handleHome).Methods("GET")
	
	fmt.Println("going to next line")
	http.ListenAndServe(":8080", r)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<H1>Welcome to Golang modules</H1>"))
}
