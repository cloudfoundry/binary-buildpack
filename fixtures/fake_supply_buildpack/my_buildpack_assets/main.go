package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(responseWriter, "Dummy App running on localhost:%s", port)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err == nil {
		log.Fatal(err)
	}
}
