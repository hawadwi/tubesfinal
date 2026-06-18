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

	err := InitDB()
	if err != nil {
		t.Fatalf("database connection failed: %v", err)
	}

	if GetDB() == nil {
		t.Fatal("database is nil")
	}

	t.Log("Gudang DB connected")
}

func setupServer() *httptest.Server {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/sort", handler.StartSort)
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_StartSort(t *testing.T) {
	setupTestDB(t)

	server := setupServer()
	defer server.Close()

	request := SortRequest{
		Resi:          "RES001",
		WarehouseZone: "Jakarta",
		Status:        "sorting",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(
		server.URL+"/sort",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}
}
