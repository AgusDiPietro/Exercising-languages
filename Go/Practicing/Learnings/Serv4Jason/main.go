// Api que retorna los datos de un usuario en formato Json cuando se ejecuta la ruta USer
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string "json:name"
	Email string "json:email"
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Agus", Email: "hola@gmail.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
