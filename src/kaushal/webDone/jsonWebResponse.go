package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	u := User{"Kaushal", 35, "BIPRD"}

	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/json", handleJson)
	http.ListenAndServe(":80", nil)
}
