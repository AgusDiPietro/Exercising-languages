package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	json.NewEncoder(w).Encode(map[string]string{
		"time": currentTime,
	})

}

func main() {
	fileServer := http.FileServer(http.Dir("./ProjetWebService"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
