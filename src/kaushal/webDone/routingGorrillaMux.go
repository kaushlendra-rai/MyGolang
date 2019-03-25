package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Install Gorilla using command
// go get -u github.com/gorilla/mux
func main() {

	r := mux.NewRouter()

	// http://localhost/books/KaushalIsAwesome/page/420
	r.HandleFunc("/books/{title}/page/{pageNo}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]
		pageNo := vars["pageNo"]

		fmt.Fprintf(w, "The requested page title: %s and pageNo: %s", title, pageNo)
	})

	http.ListenAndServe(":80", r)
}
