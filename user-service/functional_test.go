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
	dbName = "db"user
)

func TestUserFlow_Functional(t *testing.T) {

	// cet database bisa diakses
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

	// register
	email := fmt.Sprintf(
		"func%d@mail.com",
		time.Now().UnixNano(),
	)

	respReg, err := http.Post(
		"http://host.docker.internal:8081/register",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"Name":"Functional",
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

	if respReg.StatusCode != 200 {
		t.Fatalf("register failed: %+v", reg)
	}

	userID := int(reg["user_id"].(float64))

	t.Log("REGISTER SUCCESS")

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
    defer respLogin.Body.Close()

	// Gunakan map[string]interface{} agar bisa menangkap tipe data apapun (termasuk error messages)
	var login map[string]interface{}
	err = json.NewDecoder(respLogin.Body).Decode(&login)
	if err != nil {
		t.Fatalf("gagal membaca respons JSON: %v", err)
	}

	// Tampilkan pesan error aktual dari server jika status bukan 200 OK
	if respLogin.StatusCode != http.StatusOK {
		t.Fatalf("login failed with status %d. Response: %+v", respLogin.StatusCode, login)
	}

	// Pengecekan token dengan aman
	tokenInterface, exists := login["token"]
	if !exists {
		t.Fatalf("login berhasil tetapi tidak ada key 'token' di respons. Response: %+v", login)
	}

	token, ok := tokenInterface.(string)
	if !ok || token == "" {
		t.Fatalf("token kosong atau bukan string. Response: %+v", login)
	}

	t.Log("LOGIN SUCCESS")

	// get profile
	reqProfile, _ := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"http://host.docker.internal:8081/profile?id=%d",
			userID,
		),
		nil,
	)

	reqProfile.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	client := &http.Client{}

	respProfile, err := client.Do(reqProfile)

	if err != nil {
		t.Fatal(err)
	}

	var profile map[string]interface{}

	json.NewDecoder(respProfile.Body).Decode(&profile)

	if respProfile.StatusCode != 200 {
		t.Fatalf("get profile failed: %+v", profile)
	}

	t.Log("PROFILE SUCCESS")

	t.Log("FUNCTIONAL TEST SUCCESS")
}
