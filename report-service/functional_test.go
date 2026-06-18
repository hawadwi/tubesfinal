//go:build functional 
// +build functional

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFunctional_DailyReportEndpoint_ReturnsOK(t *testing.T) {

	repo := MySQLReportRepository{db: DB}
	svc := NewReportService(repo)

	h := NewReportHandler(svc)

	req := httptest.NewRequest(
		"GET",
		"/report/daily?date=2026-04-25",
		nil,
	)

	w := httptest.NewRecorder()

	h.DailyReport(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var report DailyReport
	err := json.NewDecoder(w.Body).Decode(&report)

	if err != nil {
		t.Errorf("Decode error: %v", err)
	}

	if report.TotalPaket == 0 {
		t.Errorf("Expected data > 0")
	}
}
