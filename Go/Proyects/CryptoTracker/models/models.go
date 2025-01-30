package models

import (
	"time"
)

type CoinData struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	MarketCap    float64 `json:"market_cap"`
}

type Transaction struct {
	Type        string    `json:"type"` // "buy" o "sell"
	CoinSymbol  string    `json:"coin_symbol"`
	USDAmount   float64   `json:"usd_amount"`
	TokenAmount float64   `json:"token_amount"`
	Timestamp   time.Time `json:"timestamp"`
}
