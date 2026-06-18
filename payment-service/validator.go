package main

type PaymentValidator interface {
	Validate(req PaymentRequest) bool
}

type RealPaymentValidator struct{}

func (v RealPaymentValidator) Validate(req PaymentRequest) bool {

	if req.OrderID <= 0 {
		return false
	}

	if req.Amount <= 0 {
		return false
	}

	if req.MetodePembayaran == "" {
		return false
	}

	return true
}