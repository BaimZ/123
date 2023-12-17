package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/health", healthCheckHandler)
	log.Println("server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Fprintf(w, "get people: '%v'", people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	people = append(people, person)
	fmt.Fprintf(w, "person added: '%v'", person)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "health check")
}
