package main

import (
	"errors"
	"fmt"
	"time"
)

func CalculatePayment(req CalculateRequest) *CalculateResponse {

	biayaLayanan := 0

	switch req.Layanan {

	case "reguler":
		biayaLayanan = 5000

	case "ekspres":
		biayaLayanan = 10000

	case "one-day":
		biayaLayanan = 20000

	default:
		return nil
	}

	total := (req.Berat * 1000) + (req.Jarak * 500) + biayaLayanan

	return &CalculateResponse{
		Biaya: total,
	}
}

func ProcessPayment(
	req PaymentRequest,
	v PaymentValidator,
	repo PaymentRepository,
) (Transaction, error) {

	if req.OrderID == 0 {
		return Transaction{}, errors.New("invalid order id")
	}

	if req.MetodePembayaran == "" {
		return Transaction{}, errors.New("payment method is required")
	}

	if !v.Validate(req) {
		return Transaction{}, errors.New("invalid payment request")
	}

	transaction := Transaction{
		TransactionID: fmt.Sprintf("TRX-%d", time.Now().Unix()),
		OrderID:       req.OrderID,
		Amount:        req.Amount,
		Metode:        req.MetodePembayaran,
		Status:        "SUCCESS",
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	// err := repo.Insert(transaction)
	// if err != nil {
	// 	return Transaction{}, err
	// }

	// MENJADI SEPERTI INI:
	err := repo.SaveTransaction(transaction.TransactionID, transaction.OrderID, transaction.Amount)
	if err != nil {
		return Transaction{}, err
	}

	// ===========================
	// Publish ke RabbitMQ
	// ===========================
	err = PublishPaymentEvent(PaymentEvent{
		TransactionID: transaction.TransactionID,
		OrderID:       transaction.OrderID,
		Amount:        transaction.Amount,
		Metode:        transaction.Metode,
		Status:        transaction.Status,
		Timestamp:     transaction.Timestamp,
	})

	if err != nil {
		fmt.Println("Gagal publish RabbitMQ:", err)
	}

	return transaction, nil
}

func GetTransaction(transactionID string) *Transaction {

	if transactionID == "" {
		return nil
	}

	return &Transaction{
		TransactionID: transactionID,
		OrderID:       1,
		Amount:        15000,
		Metode:        "Transfer",
		Status:        "SUCCESS",
		Timestamp:     time.Now().Format(time.RFC3339),
	}
}
