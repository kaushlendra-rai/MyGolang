package main

import (
	"fmt"
	"net/http"
)

// Use URL: http://localhost/static/
// The 'static' folder exists as a top level folder for the project.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have requested: %s, %s", r.URL.Path, r.URL.Query().Get("name"))
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
