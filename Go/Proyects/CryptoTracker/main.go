package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CoinData struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	MarketData struct {
		CurrentPrice struct {
			USD float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func getCryptoPrice(coinID string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", coinID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making the request: %v\n", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %v\n", err)
		return 0, err
	}

	var coinData CoinData
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		fmt.Printf("Error unmarshaling the JSON: %v\n", err)
		return 0, err
	}

	return coinData.MarketData.CurrentPrice.USD, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al Cripto Tracker!")
}

func priceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coinID := vars["coinID"]

	price, err := getCryptoPrice(coinID)
	if err != nil {
		fmt.Printf("Error getting crypto price: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "El precio de %s es: $%.2f USD", coinID, price)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/price/{coinID}", priceHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
		log.Fatal(err)
	}
}
