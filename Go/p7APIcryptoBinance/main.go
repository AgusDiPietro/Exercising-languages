package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Ticker struct {
	Symbol             string  `json:"symbol"`
	Price              float64 `json:"price,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	SevenDayChange     float64 `json:"sevenDayChange"`
}

func fetchTop10Cryptos() ([]Ticker, error) {
	url := "https://api.binance.com/api/v3/ticker/24hr"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tickers []Ticker
	if err := json.NewDecoder(resp.Body).Decode(&tickers); err != nil {
		return nil, err
	}

	top10 := make([]Ticker, 0, 10)
	for _, ticker := range tickers {
		if len(top10) >= 10 {
			break
		}
		if ticker.Symbol == "BTCUSDT" || ticker.Symbol == "ETHUSDT" || ticker.Symbol == "BNBUSDT" {
			sevenDayChange := calculateSevenDayChange(ticker.Symbol)
			ticker.SevenDayChange = sevenDayChange
			top10 = append(top10, ticker)
		}
	}

	return top10, nil
}

// Aquí iría una función real para calcular la variación en 7 días
func calculateSevenDayChange(symbol string) float64 {
	// Esta función debería hacer una solicitud HTTP para obtener datos históricos y calcular el cambio
	return 5.0 // Placeholder
}

func main() {
	top10, err := fetchTop10Cryptos()
	if err != nil {
		log.Fatalf("Error fetching top 10 cryptos: %v", err)
	}

	fmt.Println("Top 10 Cryptocurrencies:")
	for _, ticker := range top10 {
		fmt.Printf("Symbol: %s, Price: %.2f, 24h Change: %.2f%%, 7d Change: %.2f%%\n", ticker.Symbol, ticker.Price, ticker.PriceChangePercent, ticker.SevenDayChange)
	}
}
