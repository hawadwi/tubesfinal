//go:build functional
// +build functional
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// config
const (
	dbUser = "root"
	dbPass = "root"
	dbHost = "host.docker.internal"
	dbPort = "3306"
	dbName = "tubesdb"
)

func TestPayment_Functional(t *testing.T) {

	// cek database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatal("database tidak bisa diakses:", err)
	}

	t.Log("DATABASE CONNECTED")

	// hitung biaya
	calculateBody := []byte(`{
		"berat":2,
		"jarak":10,
		"layanan":"reguler"
	}`)

	req, _ := http.NewRequest(
		"POST",
		"http://host.docker.internal:8088/calculate",
		bytes.NewBuffer(calculateBody),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	var calculateResult map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&calculateResult)

	if resp.StatusCode != 200 {
		t.Fatalf("failed calculate: %+v", calculateResult)
	}

	t.Log("CALCULATE SUCCESS")

	// buat transaksi unik
	orderID := int(time.Now().Unix())

	paymentBody := []byte(fmt.Sprintf(`{
		"order_id": %d,
		"metode_pembayaran":"Transfer",
		"payment_details":"BCA Virtual Account",
		"amount":12000
	}`, orderID))

	req2, _ := http.NewRequest(
		"POST",
		"http://host.docker.internal:8088/pay",
		bytes.NewBuffer(paymentBody),
	)

	req2.Header.Set("Content-Type", "application/json")

	resp2, err := client.Do(req2)

	if err != nil {
		t.Fatal(err)
	}

	var payResult map[string]interface{}

	json.NewDecoder(resp2.Body).Decode(&payResult)

	if resp2.StatusCode != 200 {
		t.Fatalf("failed payment: %+v", payResult)
	}

	// cek database

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM transactions WHERE order_id = ?",
		orderID,
	).Scan(&count)

	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatal("data transaksi tidak masuk database")
	}

	t.Log("TRANSACTION SAVED")
	t.Log("FUNCTIONAL TEST SUCCESS")
}