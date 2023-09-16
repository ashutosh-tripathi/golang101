package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ashutosh-tripathi/mongoapi/controller"
	"github.com/ashutosh-tripathi/mongoapi/model"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/addMovie", addMovieHandler).Methods("POST")
	router.HandleFunc("/api/getAllMovie", getAllMovieHandler).Methods("GET")
	router.HandleFunc("/api/getMovie/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/api/updateMovie/{id}", updateMovieById).Methods("PUT")
	router.HandleFunc("/api/deleteMovie/{id}", deleteMovieById).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovie", deleteAllMovie).Methods("DELETE")

	return router

}

func addMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("movie to add:", movie)
	res := controller.InsertMovie(movie)
	w.Write([]byte("Inserted Result with id" + res))
}

func getAllMovieHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(controller.FindAllMovie())
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	id := val["id"]
	json.NewEncoder(w).Encode(controller.FindOneMovie(id))
}
func updateMovieById(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	id := val["id"]
	controller.UpdateMovie(id)
	w.Write([]byte("updated movie with id" + id))
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	controller.DeleteMovie(val["id"])
	w.Write([]byte("deleted movie with id" + val["id"]))
}

func deleteAllMovie(w http.ResponseWriter, r *http.Request) {
	controller.DeleteAllMovie()
	w.Write([]byte("deleted all"))
}
