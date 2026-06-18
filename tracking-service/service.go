package main

import "errors"

func GetTrackingStatus(resi string) *TrackingResponse {
	return nil
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
