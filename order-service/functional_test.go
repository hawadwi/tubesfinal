//go:build functional

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
	dbName = "orderdb"
)

func TestCreateOrder_Functional(t *testing.T) {

	// cek database bisa akses
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

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

	// register user
	email := fmt.Sprintf("func%d@mail.com", time.Now().UnixNano())

	respReg, err := http.Post(
		"http://host.docker.internal:8081/register",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Name":"Func",
			"Email":"%s",
			"Password":"123",
			"Role":"customer"
		}`, email))),
	)

	if err != nil {
		t.Fatal(err)
	}

	var reg map[string]interface{}
	json.NewDecoder(respReg.Body).Decode(&reg)

	userID := int(reg["user_id"].(float64))

	// login
	respLogin, err := http.Post(
		"http://host.docker.internal:8081/login",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Email":"%s",
			"Password":"123"
		}`, email))),
	)

	if err != nil {
		t.Fatal(err)
	}

	var login map[string]string
	json.NewDecoder(respLogin.Body).Decode(&login)

	token := login["token"]

	// create order
	body := []byte(fmt.Sprintf(`{
		"user_id":%d,
		"nama_barang":"Laptop",
		"berat":2,
		"dimensi":"10x10",
		"jenis":"Elektronik",
		"alamat_pengirim":"Bandung",
		"alamat_penerima":"Jakarta"
	}`, userID))

	req, _ := http.NewRequest(
		"POST",
		"http://host.docker.internal:8083/order",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != 200 {
		t.Fatalf("failed: %+v", result)
	}

	// cek data masuk database apa engga
	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE user_id = ?",
		userID,
	).Scan(&count)

	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatal("data order tidak masuk database")
	}

	t.Log("ORDER MASUK DATABASE")
	t.Log("FUNCTIONAL TEST SUCCESS")
}
