// funcion que se ejecuta antes o despues que un controlador maneje una solicitud, permite utilizarla
package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Request %s %s %s \n", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request)) {
		if r.Header.Get ("Authorization") != "Bearer token123" {
			http.Error(w, "no authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to homescreen")
}

func main() {
	http.Handle("/", authMiddleware(loggingMiddleware(http.HandleFunc(homeHandler))))
	fmt.Println("Server executing in http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
