package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func about

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server executing in http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
