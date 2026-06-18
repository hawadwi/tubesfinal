package main

type DailyReport struct {
	TotalPaket  int     `json:"total_paket"`
	Delivered   int     `json:"delivered"`
	Pending     int     `json:"pending"`
	Terlambat   int     `json:"terlambat"`
	RataRataETA float64 `json:"rata_rata_eta"`
}

type ProblemPackage struct {
	Resi   string `json:"resi"`
	Status string `json:"status"`
}

type CourierPerformance struct {
	CourierID       int     `json:"courier_id"`
	TotalPengiriman int     `json:"total_pengiriman"`
	Berhasil        int     `json:"berhasil"`
	Terlambat       int     `json:"terlambat"`
	Score           float64 `json:"score"`
}

// Struktur data dari order-service
type OrderItem struct {
	OrderID   int    `json:"order_id"`
	CreatedAt string `json:"created_at"`
}

// Struktur data dari tracking-service
type TrackingItem struct {
	Resi      string `json:"resi"`
	Lokasi    string `json:"lokasi"`
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}

// Output untuk laporan status
type StatusReport struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

// Output untuk laporan bulanan
type MonthlyReport struct {
	Month string `json:"month"`
	Total int    `json:"total"`
}
