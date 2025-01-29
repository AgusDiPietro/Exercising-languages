package main

import (
	"log"
	"net/http"

	"CryptoTracker/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Servir archivos estáticos (HTML, CSS, JS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas
	r.HandleFunc("/", handlers.HomePage)
	r.HandleFunc("/api/top10", handlers.GetTop10)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
