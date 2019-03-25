package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have requested: %s, %s", r.URL.Path, r.URL.Query().Get("name"))
	})

	http.ListenAndServe(":80", nil)
}
