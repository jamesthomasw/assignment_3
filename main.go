package main

import (
	"log"
	"net/http"

	"github.com/jamesthomasw/assignment_3/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetTankVar)

	log.Println("Connected")
	http.ListenAndServe(":8080", mux)
}
