package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Ticker struct {
	Symbol             string  `json:"symbol"`
	Price              float64 `json:"price"`
	PriceChangePercent float64 `json:"priceChangePercent"`
	SevenDayChange     float64 `json:"sevenDayChange"`
}

// Convertir los valores string a float64 en el proceso de decodificación.
type rawTicker struct {
	Symbol             string `json:"symbol"`
	Price              string `json:"lastPrice"`
	PriceChangePercent string `json:"priceChangePercent"`
}

func fetchTop10Cryptos() ([]Ticker, error) {
	url := "https://api.binance.com/api/v3/ticker/24hr"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rawTickers []rawTicker
	if err := json.NewDecoder(resp.Body).Decode(&rawTickers); err != nil {
		return nil, err
	}

	top10 := make([]Ticker, 0, 10)
	for _, rawTicker := range rawTickers {
		if len(top10) >= 10 {
			break
		}
		if rawTicker.Symbol == "BTCUSDT" || rawTicker.Symbol == "ETHUSDT" || rawTicker.Symbol == "BNBUSDT" {
			price, err := strconv.ParseFloat(rawTicker.Price, 64)
			if err != nil {
				return nil, err
			}
			priceChangePercent, err := strconv.ParseFloat(rawTicker.PriceChangePercent, 64)
			if err != nil {
				return nil, err
			}

			ticker := Ticker{
				Symbol:             rawTicker.Symbol,
				Price:              price,
				PriceChangePercent: priceChangePercent,
				SevenDayChange:     calculateSevenDayChange(rawTicker.Symbol),
			}

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
