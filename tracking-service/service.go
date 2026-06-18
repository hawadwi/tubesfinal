package main

import (
	//"database/sql"
	"errors"
)

func GetTrackingStatus(resi string) *TrackingResponse {
	// Query database untuk mencari tracking events berdasarkan resi
	query := `
	SELECT resi, lokasi, event, timestamp
	FROM tracking_events
	WHERE resi = ?
	ORDER BY id ASC
	`

	rows, err := trackingRepo.(MySQLRepository).DB.Query(query, resi)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var timeline []TrackingEvent
	for rows.Next() {
		var t TrackingEvent
		err := rows.Scan(&t.Resi, &t.Lokasi, &t.Event, &t.Timestamp)
		if err != nil {
			continue
		}
		timeline = append(timeline, t)
	}

	if len(timeline) == 0 {
		return nil
	}

	return &TrackingResponse{
		Resi:     resi,
		Status:   "in_transit",
		Timeline: timeline,
	}
}

func InsertTrackingEvent(
	req TrackingEvent,
	v ResiValidator,
	repo TrackingRepository,
) (TrackingEvent, error) {

	if !v.Validate(req.Resi) {
		return TrackingEvent{}, errors.New("invalid resi")
	}

	err := repo.Insert(req)

	if err != nil {
		return TrackingEvent{}, err
	}

	return req, nil
}

func CalculateDistance(req DistanceRequest) *DistanceResponse {
	return nil
}

func CalculateRoute(req RouteRequest) *DistanceResponse {
	return nil
}

func GetCourierLocation(courierID string) *CourierLocation {
	return nil
}
