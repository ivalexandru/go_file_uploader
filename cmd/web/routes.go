package main

import (
	"fmt"
	"log"
	"net/http"
)

func routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHome)

	fmt.Println("starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
