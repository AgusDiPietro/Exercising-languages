package main

import (
	"encoding/json"
	"fmt"
	"io"
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
		//body.StatusCode eq 403
		//	return 403
	}
	defer resp.Body.Close()

	// Lee la respuesta y lo almacena en body
	fmt.Print(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta de %s: %v", url, err)
	}

	lines := strings.Split(string(body), "\n") // Divide el cont en lineas.
	var ips []string
	for _, line := range lines {
		line = strings.TrimSpace(line)                 //remueve espacios en blanco
		if line != "" && strings.Contains(line, ".") { //busca formato IP¨"."  y luego guarda en ips
			ips = append(ips, line)
		}
	}
	return ips, nil
}

func torIPsHandler(w http.ResponseWriter, r *http.Request) {
	// URLs de las fuentes de TOR
	sources := []string{
		"https://www.dan.me.uk/torlist/?full",
		"https://torstatus.blutmagie.de",
	}

	var allIPs []TorIPResponse //Slice donde guardan las APis obtenidas
	for _, source := range sources {
		ips, err := fetchIPs(source) // Llama a una función que descarga y procesa las IPs
		if err != nil {
			fmt.Printf("Error obteniendo datos de %s: %v", source, err)
			continue
		}
		allIPs = append(allIPs, TorIPResponse{
			Source: source,
			IPs:    ips,
		})
	}

	// Codifica la respuesta en formato JSON ; w = resp from server.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allIPs); err != nil { //Utiliza el codificador JSON de Go para convertir allIPs en un formato JSON y lo escribe directamente en w
		http.Error(w, fmt.Sprintf("Error al codificar JSON: %v", err), http.StatusInternalServerError) //500
	}
}

func main() {

	// CADA 30 MIN ACTUALICE

	// Configura el endpoint
	http.HandleFunc("/tor-ips")

	// Inicia el servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Servidor iniciado en http://localhost:8080")

}
