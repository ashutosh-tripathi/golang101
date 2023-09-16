package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	AuthorId int    `json:"authorId"`
}

var courses = []Course{
	{"Course1", "MasterGo", &Author{"ashu", 123}},
	{"Course2", "MasterCICD", &Author{"bahu", 123}},
}

func main() {

	fmt.Println("starting go server...")
	router := mux.NewRouter()
	router.HandleFunc("/getCourses", getCourses).Methods("GET")
	router.HandleFunc("/getCourseByAuthor", getSpecificCourseByAuthor).Methods("GET")
	router.HandleFunc("/createCourse", addCourse).Methods("POST")
	router.HandleFunc("/updateCourse", updateCourse).Methods("POST")
	router.HandleFunc("/delCourse", deleteCourse).Methods("POST")

	http.ListenAndServe(":8080", router)
	fmt.Println("go server up and running...")

}

func getCourses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// coursebytes, err := json.MarshalIndent(courses, "", "\t")
	// if err != nil {
	// 	fmt.Println("Got error while marshalling courses")
	// }
	json.NewEncoder(w).Encode(courses)
	// w.Write(coursebytes)

}

func getSpecificCourseByAuthor(w http.ResponseWriter, r *http.Request) {
	// values := r.URL.Query()
	//another way to get url query values
	values := mux.Vars(r)
	fmt.Println("printing values", values)
	authorname := values["authorname"]
	w.Header().Set("Content-Type", "application/json")
	for _, v := range courses {
		if v.Author.Fullname == authorname {
			vbytes, _ := json.Marshal(v)
			w.Write(vbytes)
			break
		}
	}

}

func addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// bodyBytes, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("not proper data")
	// }

	var course Course

	// json.Unmarshal(bodyBytes, &course)
	//also decode using
	json.NewDecoder(r.Body).Decode(&course)
	fmt.Println("course", course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courseUpd Course
	json.NewDecoder(r.Body).Decode(&courseUpd)
	fmt.Println("printing json", courseUpd)
	for i, course := range courses {
		if courseUpd.CourseID == course.CourseID {
			fmt.Println("found couse with same id", course)
			courses[i].CourseName = courseUpd.CourseName
			courses[i].Author = courseUpd.Author
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	var delCourse []Course
	var courseUpd Course
	json.NewDecoder(r.Body).Decode(&courseUpd)
	for _, course := range courses {
		if courseUpd.CourseID != course.CourseID {
			delCourse = append(delCourse, course)
		}
	}
	courses = delCourse
}
