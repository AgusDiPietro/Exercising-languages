package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Fprintf(w, "Nombre: %s \n", name)
		fmt.Fprintf(w, "Email: %s, \n", email)
	} else {
		fmt.Fprintf(w, "Please send a form")
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server executing in http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
