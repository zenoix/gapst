package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Go API server has started")

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET / called")
		fmt.Fprint(w, "Server is running")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
