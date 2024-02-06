package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	log.Println("Starting server on port 8080...")

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

	http.Serve(listener, nil)
}
