//go:build functional 
// +build functional

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB(t *testing.T) {
	t.Helper()

	// inisialisasi koneksi database
	err := InitDB()
	if err != nil {
		t.Fatalf("database connection failed: %v", err)
	}

	// validasi apakah database berhasil terkoneksi
	if GetDB() == nil {
		t.Fatal("database is nil")
	}

	t.Log("Courier DB connected")
}

func setupServer() *httptest.Server {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	mux := http.NewServeMux()

	// endpoint untuk memulai delivery
	mux.HandleFunc("/delivery", handler.StartDelivery)

	// endpoint untuk mengambil data delivery courier
	mux.HandleFunc("/courier/deliveries", handler.GetCourierDeliveries)

	// endpoint health check service
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_StartDelivery(t *testing.T) {
	setupTestDB(t)

	server := setupServer()
	defer server.Close()

	request := DeliveryRequest{

		Resi: "RESI001",
	
		CourierID: 1,
	
		AssignedZone: "Jakarta",
	}

	// convert request ke format JSON
	body, _ := json.Marshal(request)

	resp, err := http.Post(
		server.URL+"/delivery",
		"application/json",
		bytes.NewBuffer(body),
	)

	// validasi apakah request berhasil dikirim
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	// expected response ketika delivery berhasil
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}
}
