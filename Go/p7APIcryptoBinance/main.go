package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func fetchCryptos() ([]byte, error) {
	url := "https://api.binance.com/api/v3/ticker/price"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := fetchCryptos()
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Println("Data fetched successfully")
	fmt.Println(string(data)) // Imprime los datos crudos
}
