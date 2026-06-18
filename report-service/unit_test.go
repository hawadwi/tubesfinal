package main

import (
	"context"
	"testing"
)

func TestGetDailyReport_ShouldReturnReport(t *testing.T) {

	repo := MockReportRepo{}
	svc := NewReportService(repo)

	report, err := svc.GetDailyReport(context.Background(), "2026-04-25")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if report.TotalPaket == 0 {
		t.Errorf("Expected TotalPaket > 0")
	}
}
