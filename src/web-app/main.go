package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] Received request from %s: %s %s\n", timestamp, r.RemoteAddr, r.Method, r.URL.Path)
	fmt.Fprintln(w, "Hello from Mondoo Engineer!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	fmt.Println("Server is starting...")
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
