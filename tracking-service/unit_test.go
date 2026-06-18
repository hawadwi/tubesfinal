package main

import "testing"

type MockTrackingService struct{}

func (m MockTrackingService) AddEvent(event TrackingEvent) error {
	return nil
}

func (m MockTrackingService) GetTracking(resi string) (TrackingResponse, error) {

	return TrackingResponse{
		Resi:  resi,
		Status: "Paket Diproses",
	}, nil
}

func (m MockTrackingService) GetDistance(origin, destination string) (*DistanceResponse, error) {

	return &DistanceResponse{
		DistanceKm:    120.5,
		DurationMin:   180,
		PolylineRoute: "mock-polyline",
	}, nil
}

func TestAddEvent(t *testing.T) {

	mock := MockTrackingService{}

	event := TrackingEvent{
		Resi:     "RESI001",
		Lokasi:   "Bandung",
		Event:    "Paket Diterima",
		Timestamp: "2026-05-20",
	}

	err := mock.AddEvent(event)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracking(t *testing.T) {

	mock := MockTrackingService{}

	result, err := mock.GetTracking("RESI001")

	if err != nil {
		t.Fatal(err)
	}

	if result.Resi != "RESI001" {
		t.Fail()
	}
}

func TestGetDistance(t *testing.T) {

	mock := MockTrackingService{}

	result, err := mock.GetDistance("Bandung", "Jakarta")

	if err != nil {
		t.Fatal(err)
	}

	if result.DistanceKm <= 0 {
		t.Fail()
	}
}