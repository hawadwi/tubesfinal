package main

import (
    "encoding/json"
    "net/http"
    "strings"
    "time"
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

func getAllTrackingsHandler(w http.ResponseWriter, r *http.Request) {

    repo, ok := trackingRepo.(MySQLRepository)
    if !ok {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "repository tidak valid",
        })
        return
    }

    rows, err := repo.DB.Query(`
        SELECT resi, lokasi, event, timestamp
        FROM tracking_events
        ORDER BY timestamp ASC
    `)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "error": err.Error(),
        })
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
        var ts time.Time

        err := rows.Scan(
            &t.Resi,
            &t.Lokasi,
            &t.Event,
            &ts,
        )
        if err != nil {
            continue
        }

        t.Timestamp = ts.Format(time.RFC3339)

        list = append(list, t)
    }

    if list == nil {
        list = []TrackData{}
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(list)
}