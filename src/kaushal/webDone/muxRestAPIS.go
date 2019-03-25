package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	Id      string   `json:"id,omitempty"`
	Age     int      `json:"age,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

var people []Person

func main() {
	router := mux.NewRouter()

	people = append(people, Person{"1", 30, "Sonu", &Address{"Pune", "Maharastra", "India"}})
	people = append(people, Person{"2", 31, "Rai", &Address{"Lucknow", "UP", "India"}})
	people = append(people, Person{"3", 32, "Rai", &Address{City: "Kanpur", State: "UP"}})

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":80", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// http://localhost/people/2
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person Person

	for _, p := range people {
		if p.Id == id {
			person = p
			break
		}
	}

	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
