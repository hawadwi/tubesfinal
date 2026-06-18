package main

import "testing"

type MockPaymentService struct{}

func (m MockPaymentService) Calculate(req CalculateRequest) (CalculateResponse, error) {

	biayaLayanan := 5000

	switch req.Layanan {
	case "ekspres":
		biayaLayanan = 10000
	case "one-day":
		biayaLayanan = 20000
	}

	total := (req.Berat * 1000) + (req.Jarak * 500) + biayaLayanan

	return CalculateResponse{
		Biaya: total,
	}, nil
}

func (m MockPaymentService) Pay(req PaymentRequest) (PaymentResponse, error) {

	return PaymentResponse{
		TransactionID:    "TRX001",
		StatusPembayaran: "SUCCESS",
		Biaya:            12000,
	}, nil
}

func TestCalculate(t *testing.T) {

	mock := MockPaymentService{}

	req := CalculateRequest{
		Berat:   2,
		Jarak:   10,
		Layanan: "reguler",
	}

	result, err := mock.Calculate(req)

	if err != nil {
		t.Fatal(err)
	}

	expected := (2 * 1000) + (10 * 500) + 5000

	if result.Biaya != expected {
		t.Fail()
	}
}

func TestPay(t *testing.T) {

	mock := MockPaymentService{}

	req := PaymentRequest{
		OrderID:          1,
		MetodePembayaran: "Transfer",
		PaymentDetails:   "BCA Virtual Account",
	}

	result, err := mock.Pay(req)

	if err != nil {
		t.Fatal(err)
	}

	if result.TransactionID == "" {
		t.Fail()
	}

	if result.StatusPembayaran != "SUCCESS" {
		t.Fail()
	}
}