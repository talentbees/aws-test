package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a Go web app running on AWS EC2.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
