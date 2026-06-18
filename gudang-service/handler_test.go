package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStartSort_AllValid(t *testing.T) {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	body := `{"resi":"123","warehouse_zone":"A1","status":"sorting"}`
	req := httptest.NewRequest("POST", "/sort", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler.StartSort(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestStartSort_AllErrorPaths(t *testing.T) {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	tests := []string{
		`invalid-json`,
		`{"resi":"","warehouse_zone":"A1","status":"sorting"}`,
		`{"resi":"123","warehouse_zone":"","status":"sorting"}`,
		`{"resi":"123","warehouse_zone":"A1","status":""}`,
		`{"resi":"123","warehouse_zone":"A1","status":"pending"}`,
	}

	for _, body := range tests {
		req := httptest.NewRequest("POST", "/sort", strings.NewReader(body))
		w := httptest.NewRecorder()

		handler.StartSort(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	}
}

func TestHealth_OK(t *testing.T) {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
