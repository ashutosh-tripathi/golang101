package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Phonenumber string `json:"phonenumber"`
}

var users = []User{
	{Name: "John", Email: "john@dev.com", Age: 28, Phonenumber: "12345"},
	{Name: "David", Email: "david@dev.com", Age: 30, Phonenumber: "12345"},
}

func main() {

	fmt.Println("This is a go server")

	http.HandleFunc("/", handleRoot)

	http.HandleFunc("/getUser", handleGetUser)

	http.HandleFunc("/getAllUsers", handleGetAllUsers)

	http.HandleFunc("/createUser", handleCreateUser)

	http.ListenAndServe(":8080", nil)

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root")
}
func handleGetUser(w http.ResponseWriter, r *http.Request) {

	queryvalues := r.URL.Query()
	if queryvalues.Has("name") && queryvalues.Has("age") {
		fmt.Println("checking query values for name and age", queryvalues.Get("name"), queryvalues.Get("age"))
		for _, val := range users {
			age, _ := strconv.Atoi(queryvalues.Get("age"))
			if val.Age == age && val.Name == queryvalues.Get("name") {
				fmt.Println("found user", val)
				marshalbytes, err := json.Marshal(val)
				if err != nil {
					fmt.Fprintf(w, "got error while marshalling"+err.Error())
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Println("printing marshall bytes", string(marshalbytes))
				w.Write(marshalbytes)
				return
			}
		}
	}
}
func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	userbytes, err := json.Marshal(users)
	if err != nil {
		fmt.Fprintf(w, "Got Error while getting all users"+err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userbytes)
}
func handleCreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("got request body", r.Body)
	var newuser User
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newuser)
	if err != nil {
		fmt.Println("Got error while decoding user" + err.Error())
	}
	users = append(users, newuser)

	w.Header().Set("Content-Types", "application/json")

	userbytes, err := json.Marshal(users)
	w.Write(userbytes)

}
