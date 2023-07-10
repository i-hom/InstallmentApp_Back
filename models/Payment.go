package models

type Payment struct {
	ElmakonID string `json:"elmakonid"`
	Amount    int    `json:"amount"`
	Method    string `json:"method"`
	Date      string `json:"date"`
}
