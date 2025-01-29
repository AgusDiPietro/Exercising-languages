package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"CryptoTracker/models"
)

const (
	coingeckoAPIURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
)

func GetTop10Cryptos() ([]models.CoinData, error) {
	resp, err := http.Get(coingeckoAPIURL)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	var coins []models.CoinData
	err = json.Unmarshal(body, &coins)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling the JSON: %v", err)
	}

	return coins, nil
}
