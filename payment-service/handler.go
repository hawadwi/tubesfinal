package main

import (
	"encoding/json"
	"net/http"
)

var paymentValidator PaymentValidator = RealPaymentValidator{}
var paymentRepo PaymentRepository

func calculatePaymentHandler(w http.ResponseWriter, r *http.Request) {

	var req CalculateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp := CalculatePayment(req)

	if resp == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to calculate payment",
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func processPaymentHandler(w http.ResponseWriter, r *http.Request) {

	var req PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := ProcessPayment(
		req,
		paymentValidator,
		paymentRepo,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func getTransactionHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("transaction_id")

	resp := GetTransaction(id)

	if resp == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resp)
}