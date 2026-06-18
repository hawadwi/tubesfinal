package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var resiValidator ResiValidator = RealResiValidator{}
var trackingRepo TrackingRepository

func insertTrackingEventHandler(w http.ResponseWriter, r *http.Request) {
	var req TrackingEvent
	json.NewDecoder(r.Body).Decode(&req)

	result, err := InsertTrackingEvent(
		req,
		resiValidator,
		trackingRepo,
	)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getTrackingHandler(w http.ResponseWriter, r *http.Request) {
	resi := r.URL.Query().Get("resi")
	if resi == "" {
		resi = strings.TrimPrefix(r.URL.Path, "/tracking/")
	}

	resp := GetTrackingStatus(resi)
	if resp == nil {
		// 🔥 FIXED: Berikan respons JSON agar jelas bahwa ini "Data Kosong", bukan "Salah URL"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Data tracking untuk resi " + resi + " tidak ditemukan di database",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func calculateDistanceHandler(w http.ResponseWriter, r *http.Request) {
	var req DistanceRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp := CalculateDistance(req)
	if resp == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to calculate distance",
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func calculateRouteHandler(w http.ResponseWriter, r *http.Request) {
	var req RouteRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp := CalculateRoute(req)
	if resp == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to calculate route",
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func getCourierLocationHandler(w http.ResponseWriter, r *http.Request) {
	courierID := r.URL.Query().Get("courier_id")
	if courierID == "" {
		courierID = strings.TrimPrefix(r.URL.Path, "/location/")
	}

	resp := GetCourierLocation(courierID)
	if resp == nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// 🔥 TAMBAHKAN FUNGSI INI DI PALING BAWAH FILE tracking-service/handler.go
// 🔥 VERSI PERBAIKAN: Taruh di paling bawah file tracking-service/handler.go
func getAllTrackingsHandler(w http.ResponseWriter, r *http.Request) { // 1. FIXED: Parameter diganti jadi *http.Request

	// Kita bypass pengecekan repo dengan langsung memanggil DB global milik package main jika bertipe MySQLRepository,
	// atau jika trackingRepo adalah interface, kita cast ke strukturnya.
	// Supaya aman dan pasti jalan, kita gunakan interface type assertion ke MySQLRepository:
	repo, ok := trackingRepo.(MySQLRepository)
	if !ok {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": "Repository beralih tipe atau tidak valid"})
		return
	}

	// 2. FIXED: Mengakses DB via struct repo yang sudah diekstrak
	rows, err := repo.DB.Query("SELECT resi, lokasi, event, timestamp FROM tracking_events")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	type TrackData struct {
		Resi      string `json:"resi"`
		Lokasi    string `json:"lokasi"`
		Event     string `json:"event"`
		Timestamp string `json:"timestamp"`
	}

	var list []TrackData
	for rows.Next() {
		var t TrackData
		// Scan data dari baris database
		err := rows.Scan(&t.Resi, &t.Lokasi, &t.Event, &t.Timestamp)
		if err != nil {
			continue
		}
		list = append(list, t) // FIXED: Memasukkan objek t ke dalam slice list
	}

	if list == nil {
		list = []TrackData{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
