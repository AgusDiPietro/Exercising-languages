package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AgusDiPietro/Exercising-languages/tree/main/Go/Proyects/CryptoTracker/services"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func GetTop10(w http.ResponseWriter, r *http.Request) {
	top10, err := services.GetTop10Cryptos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(top10)
}
