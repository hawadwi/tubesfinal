package main

type TrackingEvent struct {
	ID        int    `json:"id"`
	Resi      string `json:"resi"`
	Lokasi    string `json:"lokasi"`
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}

type TrackingResponse struct {
	Resi     string          `json:"resi"`
	Status   string          `json:"status"`
	Timeline []TrackingEvent `json:"timeline"`
}

type DistanceRequest struct {
	OriginAddress      string `json:"origin_address"`
	DestinationAddress string `json:"destination_address"`
}

type DistanceResponse struct {
	DistanceKm    float64 `json:"distance_km"`
	DurationMin   int     `json:"duration_min"`
	PolylineRoute string  `json:"polyline_route"`
}

type RouteRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Waypoints   []string `json:"waypoints"`
}

type CourierLocation struct {
	CourierID string  `json:"courier_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}
