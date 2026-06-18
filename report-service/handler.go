package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ReportHandler struct {
	svc *ReportService
}

func NewReportHandler(svc *ReportService) *ReportHandler {
	return &ReportHandler{svc: svc}
}

func (h *ReportHandler) DailyReport(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if date == "" {
		http.Error(w, "date required", http.StatusBadRequest)
		return
	}
	report, err := h.svc.GetDailyReport(r.Context(), date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) ProblemsReport(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := h.svc.GetProblems(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *ReportHandler) CourierPerformance(

	w http.ResponseWriter,
	r *http.Request,
) {

	id := 1
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := h.svc.GetCourierPerformance(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// 🔥 IMPROVED: Status Report Handler dengan better error handling
func (h *ReportHandler) StatusReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Validate request method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := h.svc.GetStatusReport(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// 🔥 IMPROVED: Monthly Report Handler dengan better error handling
func (h *ReportHandler) MonthlyReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Validate request method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	yearStr := r.URL.Query().Get("year")
	year := time.Now().Year()
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	data, err := h.svc.GetMonthlyReport(r.Context(), year)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
