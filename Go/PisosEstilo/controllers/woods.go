package controllers

import (
	"html/template"
	"net/http"
)

// Estructura para representar los tipos de madera
type Wood struct {
	Name        string
	Description string
	Price       float64
	ImageURL    string
}

// Handler para la página de tipos de madera
func WoodsPage(w http.ResponseWriter, r *http.Request) {
	// Ejemplo de lista de tipos de madera
	woods := []Wood{
		{"Roble", "Madera resistente y duradera.", 50.0, "/assets/img/wood1.jpg"},
		{"Pino", "Madera ligera y económica.", 30.0, "/assets/img/wood2.jpg"},
	}

	tmpl := template.Must(template.ParseFiles("templates/woods.html"))
	tmpl.Execute(w, woods)
}
