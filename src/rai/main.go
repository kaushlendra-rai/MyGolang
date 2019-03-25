package main

import (
	"log"
	"net/http"
	"rai/handler"
)

func main() {
	router := handler.Initialize()
	log.Fatal(http.ListenAndServe(":80", router))
}
