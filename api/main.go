package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type predictionRequest struct {
	SepalLengthCm float32
	SepalWidthCm  float32
	PetalLengthCm float32
	PetalWidthCm  float32
}

func main() {
	mux := http.NewServeMux()

	log.Println("Go API server has started")

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET / called")
		fmt.Fprint(w, "Server is running")
	})

	mux.HandleFunc("POST /predict", func(w http.ResponseWriter, r *http.Request) {
		log.Println("POST /predict called")

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var p predictionRequest

		if err := d.Decode(&p); err != nil {
			log.Printf("Error occurred while decoding JSON object - %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if d.More() {
			log.Println("Extraneous data found after JSON object")
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
			return
		}

		log.Println(p)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
