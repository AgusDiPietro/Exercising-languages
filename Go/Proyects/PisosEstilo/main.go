package main

import (
	"PisosEstilo/controllers"
	"html/template"
	"log"
	"net/http"
)

// Estructura básica para datos que vas a mostrar en las páginas
type Work struct {
	Title       string
	Description string
	ImageURL    string
}

func main() {
	// Servir archivos estáticos desde la carpeta assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas a secciones
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/jobs", controllers.JobsPage)
	http.HandleFunc("/woods", controllers.WoodsPage)
	http.HandleFunc("/quote", controllers.QuotePage)

	// Iniciar el servidor
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Funciones para cada página
func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}
