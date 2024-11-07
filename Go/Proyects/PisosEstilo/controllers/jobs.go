package controllers

import (
	"html/template"
	"net/http"
)

// Estructura para representar los trabajos
type Job struct {
	Title       string
	Description string
	ImageURL    string
}

// Handler para la página de trabajos
func JobsPage(w http.ResponseWriter, r *http.Request) {
	// Ejemplo de lista de trabajos
	jobs := []Job{
		{"Instalación de Pisos de Madera", "Instalación de alta calidad con acabado profesional.", "/assets/img/job1.jpg"},
		{"Restauración de Pisos", "Renovamos tus pisos para que luzcan como nuevos.", "/assets/img/job2.jpg"},
	}

	tmpl := template.Must(template.ParseFiles("templates/jobs.html"))
	tmpl.Execute(w, jobs)
}
