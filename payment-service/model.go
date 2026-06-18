package main

type CalculateRequest struct {
	Berat   int    `json:"berat"`
	Jarak   int    `json:"jarak"`
	Layanan string `json:"layanan"`
}

type CalculateResponse struct {
	Biaya int `json:"biaya"`
}

type PaymentRequest struct {
	OrderID           int    `json:"order_id"`
	MetodePembayaran  string `json:"metode_pembayaran"`
	PaymentDetails    string `json:"payment_details"`
	Amount             int    `json:"amount"`
}

type PaymentResponse struct {
	TransactionID   string `json:"transaction_id"`
	StatusPembayaran string `json:"status_pembayaran"`
	Biaya           int    `json:"biaya"`
}

type PaymentEvent struct {
	TransactionID string  `json:"transaction_id"`
	OrderID       int     `json:"order_id"`
	Amount        int    `json:"amount"`
	Metode        string  `json:"metode"`
	Status        string  `json:"status"`
	Timestamp     string  `json:"timestamp"`
}

type Transaction struct {
	TransactionID string `json:"transaction_id"`
	OrderID       int    `json:"order_id"`
	Amount        int    `json:"amount"`
	Metode        string `json:"metode"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}