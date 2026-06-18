package main

import (
	"database/sql"
	"time"
)

type MySQLRepository struct {
	DB *sql.DB
}

func (r MySQLRepository) SaveTransaction(txID string, orderID int, amount int) error {
	query := `
	INSERT INTO transactions (transaction_id, order_id, amount, metode, status, timestamp)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	// Saat pertama kali masuk dari queue, kita set metode "BELUM_DISELEKSI" dan status "pending"
	_, err := r.DB.Exec(query, txID, orderID, amount, "PENDING_SELECTION", "pending", time.Now())
	return err
}
