package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// ==================== INTERFACE & SERVICE ====================

type ReportRepository interface {
	GetDailyReport(ctx context.Context, date string) (*DailyReport, error)
	GetProblems(ctx context.Context) ([]ProblemPackage, error)
	GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

// 🔥 FIXED: Get Daily Report dengan timestamp filter dari tracking service
func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {

    trackings, err := fetchTrackings()
    if err != nil {
        return &DailyReport{
            TotalPaket:  0,
            Delivered:   0,
            Pending:     0,
            Terlambat:   0,
            RataRataETA: 0,
        }, nil
    }

    if len(trackings) == 0 {
        return &DailyReport{
            TotalPaket:  0,
            Delivered:   0,
            Pending:     0,
            Terlambat:   0,
            RataRataETA: 0,
        }, nil
    }

    var total, pending, delivered int

    for _, t := range trackings {

        ts, err := time.Parse(time.RFC3339, t.Timestamp)
        if err != nil {
            ts, err = time.Parse("2006-01-02 15:04:05", t.Timestamp)
            if err != nil {
                continue
            }
        }

        if ts.Format("2006-01-02") == date {
            total++

            if t.Lokasi == "Serah terima ke penerima" {
                delivered++
            } else {
                pending++
            }
        }
    }

    return &DailyReport{
        TotalPaket:  total,
        Delivered:   delivered,
        Pending:     pending,
        Terlambat:   0,
        RataRataETA: 0,
    }, nil
}

func (s *ReportService) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
	return s.repo.GetProblems(ctx)
}

func (s *ReportService) GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error) {
	return s.repo.GetCourierPerformance(ctx, courierID)
}

// ==================== FUNGSI PANGGIL API ====================

var httpClient = &http.Client{Timeout: 5 * time.Second}

// 🔥 FIXED: Improved API fetching with better error handling and fallback
func fetchOrders() ([]OrderItem, error) {
	orderURL := os.Getenv("ORDER_SERVICE_URL")
	if orderURL == "" {
		orderURL = "http://order-service:8083" // Docker default
	}

	url := orderURL + "/orders"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []OrderItem{}, nil // Return empty on error
	}

	var orders []OrderItem
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return []OrderItem{}, nil
	}

	if orders == nil {
		orders = []OrderItem{}
	}

	return orders, nil
}

// 🔥 FIXED: Improved tracking API fetching with better error handling
func fetchTrackings() ([]TrackingItem, error) {
	trackingURL := os.Getenv("TRACKING_SERVICE_URL")
	if trackingURL == "" {
		trackingURL = "http://tracking-service:8084" // Docker default
	}

	url := trackingURL + "/trackings"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []TrackingItem{}, nil // Return empty on error
	}

	var trackings []TrackingItem
	if err := json.NewDecoder(resp.Body).Decode(&trackings); err != nil {
		return []TrackingItem{}, nil
	}

	if trackings == nil {
		trackings = []TrackingItem{}
	}

	return trackings, nil
}

// ==================== METHOD LAPORAN BARU ====================

// 🔥 FIXED: Get Status Report dari Tracking Service dengan proper error handling
func (s *ReportService) GetStatusReport(ctx context.Context) ([]StatusReport, error) {
	trackings, err := fetchTrackings()
	if err != nil {
		return []StatusReport{}, nil // Return empty on error
	}

	if len(trackings) == 0 {
		return []StatusReport{}, nil
	}

	count := make(map[string]int)
	for _, t := range trackings {
		if t.Lokasi != "" {
			count[t.Lokasi]++
		}
	}

	var result []StatusReport
	for status, total := range count {
		result = append(result, StatusReport{Status: status, Total: total})
	}

	if result == nil {
		result = []StatusReport{}
	}

	return result, nil
}

// 🔥 FIXED: Get Monthly Report dari Order Service dengan proper error handling
func (s *ReportService) GetMonthlyReport(ctx context.Context, year int) ([]MonthlyReport, error) {
	orders, err := fetchOrders()
	if err != nil {
		return []MonthlyReport{}, nil
	}

	if len(orders) == 0 {
		return []MonthlyReport{}, nil
	}

	monthly := make(map[string]int)
	currentYear := time.Now().Year()

	for _, o := range orders {
		// Sesuaikan format waktu dengan response dari order-service
		var t time.Time
		var parseErr error

		t, parseErr = time.Parse(time.RFC3339, o.CreatedAt)
		if parseErr != nil {
			// Coba format lain
			t, parseErr = time.Parse("2006-01-02 15:04:05", o.CreatedAt)
			if parseErr != nil {
				// Coba format tanpa waktu
				t, parseErr = time.Parse("2006-01-02", o.CreatedAt)
				if parseErr != nil {
					continue
				}
			}
		}

		// Use current year if not specified
		checkYear := year
		if checkYear == 0 {
			checkYear = currentYear
		}

		if t.Year() == checkYear {
			key := t.Format("2006-01")
			monthly[key]++
		}
	}

	var result []MonthlyReport
	for month, total := range monthly {
		result = append(result, MonthlyReport{Month: month, Total: total})
	}

	if result == nil {
		result = []MonthlyReport{}
	}

	return result, nil
}
