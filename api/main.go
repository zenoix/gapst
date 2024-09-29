package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/zenoix/gapst/pb"
)

type predictionRequest struct {
	SepalLengthCm float32
	SepalWidthCm  float32
	PetalLengthCm float32
	PetalWidthCm  float32
}

func main() {
	serverAddr := flag.String(
		"server", "localhost:50051",
		"The server address in the format of host:port",
	)
	flag.Parse()

	mux := http.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}
	defer conn.Close()

	client := pb.NewModelClient(conn)

	log.Println("gRPC connection made successfully")

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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := client.Predict(ctx, &pb.PredictRequest{

			SepalLengthCm: p.SepalLengthCm,
			SepalWidthCm:  p.SepalWidthCm,
			PetalLengthCm: p.PetalLengthCm,
			PetalWidthCm:  p.PetalWidthCm,
		})
		if err != nil {
			log.Println("Error sending request:", err)
			http.Error(w, "error sending request", http.StatusInternalServerError)
		}

		log.Println(res.Species)
		fmt.Fprintf(w, res.Species)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
