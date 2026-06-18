// package main

// import (
// 	"database/sql"
// )

// type PackageRepository struct {
// 	db *sql.DB
// }

// func NewPackageRepository(db *sql.DB) *PackageRepository {
// 	return &PackageRepository{db: db}
// }

// // ======================
// // CREATE PACKAGE (FIXED SINKRONISASI KOLOM & ?)
// // ======================
// func (r *PackageRepository) Create(pkg *Package) error {

// 	_, err := r.db.Exec(
// 		`INSERT INTO packages
// 		(user_id,
// 		 resi,
// 		 nama_barang,
// 		 berat,
// 		 dimensi,
// 		 jenis,
// 		 alamat_pengirim,
// 		 alamat_penerima,
// 		 nama_penerima,
// 		 no_telp_penerima,
// 		 warehouse_zone,
// 		 status)
// 		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, // 🔥 SEKARANG PAS 12 TANDA TANYA
// 		pkg.UserID,
// 		pkg.Resi,
// 		pkg.NamaBarang,
// 		pkg.Berat,
// 		pkg.Dimensi,
// 		pkg.Jenis,
// 		pkg.AlamatPengirim,
// 		pkg.AlamatPenerima,
// 		pkg.NamaPenerima,
// 		pkg.NoTelpPenerima,
// 		pkg.WarehouseZone,
// 		pkg.Status,
// 	)

// 	return err
// }

// // ======================
// // GET BY RESI
// // ======================
// func (r *PackageRepository) GetByResi(resi string) (*Package, error) {

// 	var pkg Package

// 	err := r.db.QueryRow(
// 		`SELECT
// 			user_id,
// 			resi,
// 			nama_barang,
// 			berat,
// 			dimensi,
// 			jenis,
// 			alamat_pengirim,
// 			alamat_penerima,
// 			nama_penerima,
// 			no_telp_penerima,
// 			status,
// 			warehouse_zone,
// 			created_at,
// 			sorted_at
// 		FROM packages
// 		WHERE resi = ?`,
// 		resi,
// 	).Scan(
// 		&pkg.UserID,
// 		&pkg.Resi,
// 		&pkg.NamaBarang,
// 		&pkg.Berat,
// 		&pkg.Dimensi,
// 		&pkg.Jenis,
// 		&pkg.AlamatPengirim,
// 		&pkg.AlamatPenerima,
// 		&pkg.NamaPenerima,
// 		&pkg.NoTelpPenerima,
// 		&pkg.Status,
// 		&pkg.WarehouseZone,
// 		&pkg.CreatedAt,
// 		&pkg.SortedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pkg, nil
// }

// // COMPLETE SORT
// func (r *PackageRepository) CompleteSort(resi string) error {

// 	_, err := r.db.Exec(
// 		`UPDATE packages
// 		SET status='ready',
// 			sorted_at=NOW()
// 		WHERE resi=?`,
// 		resi,
// 	)

// 	return err
// }

// // OUTBOX
// func (r *PackageRepository) SaveOutbox(eventType string, payload string) error {
// 	_, err := r.db.Exec(
// 		`INSERT INTO outbox_events (event_type, payload, status)
// 		 VALUES (?, ?, 'pending')`,
// 		eventType,
// 		payload,
// 	)
// 	return err
// }

// // UPDATE STATUS
// func (r *PackageRepository) UpdateStatus(resi string, status string) error {
// 	_, err := r.db.Exec(
// 		`UPDATE packages SET status=? WHERE resi=?`,
// 		status,
// 		resi,
// 	)
// 	return err
// }

// // OUTBOX FETCH
// func (r *PackageRepository) GetPendingEvents() ([]OutboxEvent, error) {

// 	rows, err := r.db.Query(
// 		`SELECT id, event_type, payload
// 		 FROM outbox_events
// 		 WHERE status='pending'`,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var events []OutboxEvent

// 	for rows.Next() {
// 		var e OutboxEvent
// 		rows.Scan(&e.ID, &e.EventType, &e.Payload)
// 		events = append(events, e)
// 	}

// 	return events, nil
// }

// // MARK SENT
// func (r *PackageRepository) MarkAsSent(id int) error {
// 	_, err := r.db.Exec(
// 		`UPDATE outbox_events
// 		 SET status='sent', sent_at=NOW()
// 		 WHERE id=?`,
// 		id,
// 	)
// 	return err
// }

// // ======================
// // GET ALL (FIXED SCAN ORDER SINKRON)
// // ======================
// func (r *PackageRepository) GetAll() ([]Package, error) {

// 	rows, err := r.db.Query(`
// 		SELECT
// 			user_id,
// 			resi,
// 			nama_barang,
// 			berat,
// 			dimensi,
// 			jenis,
// 			alamat_pengirim,
// 			alamat_penerima,
// 			nama_penerima,
// 			no_telp_penerima,
// 			status,
// 			warehouse_zone,
// 			created_at,
// 			sorted_at
// 		FROM packages
// 		ORDER BY resi
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var packages []Package

// 	for rows.Next() {

// 		var p Package

// 		// 🔥 FIXED: Scan diurutkan persis sesuai susunan SELECT di atas agar tidak bergeser kolomnya
// 		err := rows.Scan(
// 			&p.UserID,
// 			&p.Resi,
// 			&p.NamaBarang,
// 			&p.Berat,
// 			&p.Dimensi,
// 			&p.Jenis,
// 			&p.AlamatPengirim,
// 			&p.AlamatPenerima,
// 			&p.NamaPenerima,
// 			&p.NoTelpPenerima,
// 			&p.Status,
// 			&p.WarehouseZone,
// 			&p.CreatedAt,
// 			&p.SortedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		packages = append(packages, p)
// 	}

// 	// Cegah kembalian JSON bernilai null jika tabel kosong
// 	if packages == nil {
// 		packages = []Package{}
// 	}

// 	return packages, nil
// }

// // Taruh di bagian service/handler gudang saat proses sortir selesai dilakukan
// func PublishGudangEventToTracking(rabbitChannel *amqp.Channel, resi string, zone string) {
// 	event := map[string]interface{}{
// 		"resi":      resi,
// 		"lokasi":    "Gudang Zone " + zone,
// 		"event":     "Paket Selesai Disortir - Siap Dikirim",
// 		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
// 	}
// 	body, _ := json.Marshal(event)

// 	// Kirim langsung ke tracking_queue
// 	rabbitChannel.Publish(
// 		"",               // exchange
// 		"tracking_queue", // routing key / nama queue tracking
// 		false, false,
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		},
// 	)
// }

package main

import (
	"database/sql"
	"encoding/json" // 🔥 FIXED: Ditambahkan
	"time"          // 🔥 FIXED: Ditambahkan

	"github.com/streadway/amqp" // 🔥 FIXED: Ditambahkan
)

type PackageRepository struct {
	db *sql.DB
}

func NewPackageRepository(db *sql.DB) *PackageRepository {
	return &PackageRepository{db: db}
}

// ======================
// CREATE PACKAGE
// ======================
func (r *PackageRepository) Create(pkg *Package) error {
	_, err := r.db.Exec(
		`INSERT INTO packages
		(user_id, resi, nama_barang, berat, dimensi, jenis, alamat_pengirim, alamat_penerima, nama_penerima, no_telp_penerima, warehouse_zone, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		pkg.UserID,
		pkg.Resi,
		pkg.NamaBarang,
		pkg.Berat,
		pkg.Dimensi,
		pkg.Jenis,
		pkg.AlamatPengirim,
		pkg.AlamatPenerima,
		pkg.NamaPenerima,
		pkg.NoTelpPenerima,
		pkg.WarehouseZone,
		pkg.Status,
	)
	return err
}

// ======================
// GET BY RESI
// ======================
func (r *PackageRepository) GetByResi(resi string) (*Package, error) {
	var pkg Package
	err := r.db.QueryRow(
		`SELECT user_id, resi, nama_barang, berat, dimensi, jenis, alamat_pengirim, alamat_penerima, nama_penerima, no_telp_penerima, status, warehouse_zone, created_at, sorted_at
		FROM packages WHERE resi = ?`,
		resi,
	).Scan(
		&pkg.UserID, &pkg.Resi, &pkg.NamaBarang, &pkg.Berat, &pkg.Dimensi, &pkg.Jenis, &pkg.AlamatPengirim, &pkg.AlamatPenerima, &pkg.NamaPenerima, &pkg.NoTelpPenerima, &pkg.Status, &pkg.WarehouseZone, &pkg.CreatedAt, &pkg.SortedAt,
	)
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

// COMPLETE SORT
func (r *PackageRepository) CompleteSort(resi string) error {
	_, err := r.db.Exec(
		`UPDATE packages SET status='ready', sorted_at=NOW() WHERE resi=?`,
		resi,
	)
	return err
}

// OUTBOX
func (r *PackageRepository) SaveOutbox(eventType string, payload string) error {
	_, err := r.db.Exec(
		`INSERT INTO outbox_events (event_type, payload, status) VALUES (?, ?, 'pending')`,
		eventType, payload,
	)
	return err
}

// UPDATE STATUS
func (r *PackageRepository) UpdateStatus(resi string, status string) error {
	_, err := r.db.Exec(
		`UPDATE packages SET status=? WHERE resi=?`,
		status, resi,
	)
	return err
}

// OUTBOX FETCH
func (r *PackageRepository) GetPendingEvents() ([]OutboxEvent, error) {
	rows, err := r.db.Query(`SELECT id, event_type, payload FROM outbox_events WHERE status='pending'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []OutboxEvent
	for rows.Next() {
		var e OutboxEvent
		rows.Scan(&e.ID, &e.EventType, &e.Payload)
		events = append(events, e)
	}
	return events, nil
}

// MARK SENT
func (r *PackageRepository) MarkAsSent(id int) error {
	_, err := r.db.Exec(`UPDATE outbox_events SET status='sent', sent_at=NOW() WHERE id=?`, id)
	return err
}

// ======================
// GET ALL
// ======================
func (r *PackageRepository) GetAll() ([]Package, error) {
	rows, err := r.db.Query(`
		SELECT user_id, resi, nama_barang, berat, dimensi, jenis, alamat_pengirim, alamat_penerima, nama_penerima, no_telp_penerima, status, warehouse_zone, created_at, sorted_at
		FROM packages ORDER BY resi
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []Package
	for rows.Next() {
		var p Package
		err := rows.Scan(
			&p.UserID, &p.Resi, &p.NamaBarang, &p.Berat, &p.Dimensi, &p.Jenis, &p.AlamatPengirim, &p.AlamatPenerima, &p.NamaPenerima, &p.NoTelpPenerima, &p.Status, &p.WarehouseZone, &p.CreatedAt, &p.SortedAt,
		)
		if err != nil {
			return nil, err
		}
		packages = append(packages, p)
	}

	if packages == nil {
		packages = []Package{}
	}
	return packages, nil
}

// 🔥 FIXED: Ditambahkan fungsi pembantu pengiriman event ke tracking-service
func PublishGudangEventToTracking(rabbitChannel *amqp.Channel, resi string, zone string) {
	event := map[string]interface{}{
		"resi":      resi,
		"lokasi":    "Gudang Zone " + zone,
		"event":     "Paket Selesai Disortir - Siap Dikirim",
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}
	body, _ := json.Marshal(event)

	if rabbitChannel != nil {
		_ = rabbitChannel.Publish(
			"",
			"tracking_queue",
			false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
	}
}
