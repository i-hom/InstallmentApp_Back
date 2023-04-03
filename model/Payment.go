package model

type Payment struct {
	InstallmentID string `json:"installment_id"`
	Amount        int    `json:"amount"`
	Method        string `json:"payment_method"`
	Date          string `json:"date"`
}
