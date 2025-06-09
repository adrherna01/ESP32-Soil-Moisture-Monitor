package main

import (
	"log"
	"net/http"
)

type Measurement struct {
	Temperature float64 `json:"temperature"`
	Humidity float64 `json:"humidity"`
	Timestamp float64 `json:"timestamp"`
}

func main() {
	http.HandleFunc("/measurements", measurementHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func measurementHandler() {
	
}