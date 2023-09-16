package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"Name"`
	Price    float64  `json:"price"`
	Password string   `json:"-"`
	Tech     []string `json:"tech,omitempty"`
}

var courses = []Course{
	{"Golang", 100, "123", []string{"go", "postman"}},
	{"java", 200.200, "456", []string{"java", "spring", "springboot"}},
	{"hlf", 1000.1, "789", nil},
}

func main() {

	marjson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		fmt.Println("got error while marshalling courses" + err.Error())
	}

	// fmt.Println("marjson", marjson)
	// fmt.Println("marjson to string", string(marjson))

	//unmarshalling json

	var coursesjson []Course
	if json.Valid(marjson) {
		err := json.Unmarshal(marjson, &coursesjson)
		if err != nil {
			fmt.Println("got error while unmarshalling courses" + err.Error())
		}
		fmt.Printf("%#v", coursesjson)
	} else {
		fmt.Println("not valid json")
	}
	//json as map
	var mapCourse []map[string]interface{}
	err = json.Unmarshal(marjson, &mapCourse)
	if err != nil {
		fmt.Println("Got error while unmarshalling")
	}
	fmt.Printf("%+v\n", mapCourse)
	fmt.Println("map value" + mapCourse[1]["Name"].(string))

}
