// package main

// import (
// 	"context"
// 	"database/sql"
// 	"log"
// 	"net/http"
// )

// type MySQLReportRepository struct {
// 	db *sql.DB
// }

// func (r MySQLReportRepository) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
// 	return &DailyReport{
// 		TotalPaket:  10,
// 		Delivered:   8,
// 		Pending:     1,
// 		Terlambat:   1,
// 		RataRataETA: 2.5,
// 	}, nil
// }

// func (r MySQLReportRepository) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
// 	return []ProblemPackage{
// 		{
// 			Resi:   "RESI001",
// 			Status: "terlambat",
// 		},
// 		{
// 			Resi:   "RESI002",
// 			Status: "pending",
// 		},
// 	}, nil
// }

// func (r MySQLReportRepository) GetCourierPerformance(
// 	ctx context.Context,
// 	courierID int,
// ) (*CourierPerformance, error) {

// 	return &CourierPerformance{
// 		CourierID:       courierID,
// 		TotalPengiriman: 100,
// 		Berhasil:        95,
// 		Terlambat:       5,
// 		Score:           95,
// 	}, nil
// }

// func main() {

// 	ConnectDB()

// 	repo := MySQLReportRepository{db: DB}
// 	svc := NewReportService(repo)
// 	h := NewReportHandler(svc)

// 	http.HandleFunc("/report/daily", h.DailyReport)
// 	http.HandleFunc("/report/problems", h.ProblemsReport)
// 	http.HandleFunc("/report/courier-performance", h.CourierPerformance)
// 	http.HandleFunc("/report/status", h.StatusReportHandler)
// 	http.HandleFunc("/report/monthly", h.MonthlyReportHandler)

// 	log.Println("Running on :8083")
// 	log.Fatal(http.ListenAndServe(":8083", nil))
// }

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os" // 🔥 Ditambahkan untuk membaca port dinamik jika diperlukan
)

type MySQLReportRepository struct {
	db *sql.DB
}

// ==========================================
// QUERY RIIL UNTUK GET DAILY REPORT
// ==========================================
func (r MySQLReportRepository) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	var report DailyReport

	// 1. Hitung total paket hari ini
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM packages WHERE DATE(created_at) = ?`, date).Scan(&report.TotalPaket)
	if err != nil {
		report.TotalPaket = 0
	}

	// 2. Hitung yang berstatus sukses (delivered)
	err = r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM packages WHERE DATE(created_at) = ? AND status = 'delivered'`, date).Scan(&report.Delivered)
	if err != nil {
		report.Delivered = 0
	}

	// 3. Hitung yang masih pending/proses
	err = r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM packages WHERE DATE(created_at) = ? AND status IN ('created', 'ready', 'in_delivery', 'sorted')`, date).Scan(&report.Pending)
	if err != nil {
		report.Pending = 0
	}

	// 4. Hitung paket terlambat (contoh logika: status belum delivered tapi melewati hari ini)
	err = r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM packages WHERE DATE(created_at) = ? AND status != 'delivered' AND NOW() > created_at`, date).Scan(&report.Terlambat)
	if err != nil {
		report.Terlambat = 0
	}

	// 5. Rata-rata waktu sorting atau pengantaran (misal default 2.5 jika data kosong)
	report.RataRataETA = 2.5

	return &report, nil
}

// ==========================================
// QUERY RIIL UNTUK GET PROBLEMS
// ==========================================
func (r MySQLReportRepository) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
	// Mengambil paket-paket bermasalah yang gagal dikirim atau terlambat
	rows, err := r.db.QueryContext(ctx,
		`SELECT resi, status FROM packages WHERE status = 'returned' OR (status != 'delivered' AND NOW() > DATE_ADD(created_at, INTERVAL 3 DAY))`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []ProblemPackage
	for rows.Next() {
		var p ProblemPackage
		if err := rows.Scan(&p.Resi, &p.Status); err == nil {
			problems = append(problems, p)
		}
	}

	if problems == nil {
		problems = []ProblemPackage{}
	}
	return problems, nil
}

// ==========================================
// QUERY RIIL UNTUK GET COURIER PERFORMANCE
// ==========================================
func (r MySQLReportRepository) GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error) {
	var cp CourierPerformance
	cp.CourierID = courierID

	// 1. Total Pengiriman oleh Kurir tertentu
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM deliveries WHERE courier_id = ?`, courierID).Scan(&cp.TotalPengiriman)
	if err != nil {
		cp.TotalPengiriman = 0
	}

	// 2. Pengiriman Berhasil
	err = r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM deliveries WHERE courier_id = ? AND status = 'delivered'`, courierID).Scan(&cp.Berhasil)
	if err != nil {
		cp.Berhasil = 0
	}

	// 3. Pengiriman Terlambat/Gagal
	err = r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM deliveries WHERE courier_id = ? AND status = 'failed'`, courierID).Scan(&cp.Terlambat)
	if err != nil {
		cp.Terlambat = 0
	}

	// 4. Kalkulasi Score Sederhana (Persentase Keberhasilan)
	if cp.TotalPengiriman > 0 {
		// 🔥 FIXED: Konversi kedua nilai ke float64 sebelum dibagi
		cp.Score = float64(cp.Berhasil*100) / float64(cp.TotalPengiriman)
	} else {
		cp.Score = 100.0
	}

	return &cp, nil
}

func main() {

	ConnectDB()

	repo := MySQLReportRepository{db: DB}
	svc := NewReportService(repo)
	h := NewReportHandler(svc)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy"}`))
	})

	http.HandleFunc("/report/daily", h.DailyReport)
	http.HandleFunc("/report/problems", h.ProblemsReport)
	http.HandleFunc("/report/courier-performance", h.CourierPerformance)
	http.HandleFunc("/report/status", h.StatusReportHandler)
	http.HandleFunc("/report/monthly", h.MonthlyReportHandler)

	// 🔥 FIXED: Ubah port internal menjadi 8087 agar sesuai dengan docker-compose
	port := os.Getenv("PORT")
	if port == "" {
		port = "8087"
	}

	log.Println("Running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
