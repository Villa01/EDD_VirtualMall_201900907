package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	router := CreateRouter()
	log.Fatal(http.ListenAndServe(port, router))

}
