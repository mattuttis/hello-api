package main

import (
	"log"
	"net/http"

	"github.com/mattuttis/hello-api/handlers/rest"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello/", rest.TranslateHandler)

	log.Printf("Listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
