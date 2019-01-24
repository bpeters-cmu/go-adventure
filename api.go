package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllJobsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateJobEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateJobEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteJobEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindJobEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/jobs", AllJobsEndPoint).Methods("GET")
	r.HandleFunc("/jobs", CreateJobEndPoint).Methods("POST")
	r.HandleFunc("/jobs", UpdateJobEndPoint).Methods("PUT")
	r.HandleFunc("/jobs", DeleteJobEndPoint).Methods("DELETE")
	r.HandleFunc("/jobs/{id}", FindJobEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
