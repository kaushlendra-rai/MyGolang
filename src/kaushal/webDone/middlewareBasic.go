package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before: ", r.URL.Path)
		f(w, r)
		log.Println("Done for: ", r.URL.Path)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "From Foo")
	log.Println("Foo() invoked")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "From Bar")
	log.Println("Bar() invoked")
}

func main() {
	http.HandleFunc("/foo", handle(foo))
	http.HandleFunc("/bar", handle(bar))

	http.ListenAndServe(":80", nil)
}
