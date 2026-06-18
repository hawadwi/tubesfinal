package main

type PaymentRepository interface {
	// Pastikan ada cetak biru method ini (sesuaikan dengan kode yang sudah kamu punya)
	SaveTransaction(txID string, orderID int, amount int) error
}
