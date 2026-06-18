package main

import (
	"context"
)

type MockReportRepo struct{}

func (m MockReportRepo) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}

func (m MockReportRepo) GetProblems(ctx context.Context) ([]ProblemPackage, error) {

	return []ProblemPackage{
		{
			Resi:   "RESI001",
			Status: "terlambat",
		},
		{
			Resi:   "RESI002",
			Status: "pending",
		},
	}, nil
}

func (m MockReportRepo) GetCourierPerformance(
	ctx context.Context,
	courierID int,
) (*CourierPerformance, error) {

	return &CourierPerformance{
		CourierID:       courierID,
		TotalPengiriman: 100,
		Berhasil:        95,
		Terlambat:       5,
		Score:           95,
	}, nil
}
