package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TorIPResponse struct {
	Source string   `json:"source"`
	IPs    []string `json:"ips"`
}

func fetchIPs(url string) ([]string, error) {
	// Realiza una solicitud GET al URL proporcionado
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener datos de %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Lee la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta de %s: %v", url, err)
	}

	// Divide las líneas para extraer las IPs
	lines := strings.Split(string(body), "\n")
	var ips []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && strings.Contains(line, ".") {
			ips = append(ips, line)
		}
	}
	return ips, nil
}

func torIPsHandler(w http.ResponseWriter, r *http.Request) {
	// URLs de las fuentes de TOR
	sources := []string{
		"https://www.dan.me.uk/tornodes",
		"https://torstatus.blutmagie.de",
	}

	var allIPs []TorIPResponse
	for _, source := range sources {
		ips, err := fetchIPs(source)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error obteniendo datos de %s: %v", source, err), http.StatusInternalServerError)
			return
		}
		allIPs = append(allIPs, TorIPResponse{
			Source: source,
			IPs:    ips,
		})
	}

	// Codifica la respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allIPs); err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar JSON: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	// Configura el endpoint
	http.HandleFunc("/tor-ips", torIPsHandler)

	// Inicia el servidor
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
