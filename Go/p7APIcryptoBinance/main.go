package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para almacenar datos de la respuesta de Binance
type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func fetchCryptos() ([]Ticker, error) {
	url := "https://api.binance.com/api/v3/ticker/price"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tickers []Ticker
	if err := json.NewDecoder(resp.Body).Decode(&tickers); err != nil {
		return nil, err
	}

	return tickers, nil
}

func main() {
	tickers, err := fetchCryptos()
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Println("Data fetched and decoded successfully")
	for _, ticker := range tickers {
		fmt.Printf("Symbol: %s, Price: %s\n", ticker.Symbol, ticker.Price)
	}
}
