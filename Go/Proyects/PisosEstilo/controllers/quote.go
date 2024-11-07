package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type QuoteRequest struct {
	Name    string
	Email   string
	Phone   string
	Message string
}

// Renderiza la página de solicitud de presupuesto
func QuotePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/quote.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// Capturar los datos del formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error procesando el formulario", http.StatusInternalServerError)
			return
		}

		// Crear una instancia de QuoteRequest con los datos del formulario
		quote := QuoteRequest{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Phone:   r.FormValue("phone"),
			Message: r.FormValue("message"),
		}

		// Procesar la solicitud o enviar por email (esto es solo un log para ejemplo)
		log.Printf("Solicitud de presupuesto recibida: %+v\n", quote)

		// Enviar una respuesta de confirmación
		tmpl := template.Must(template.ParseFiles("templates/quote_success.html"))
		tmpl.Execute(w, quote)
	}
}
