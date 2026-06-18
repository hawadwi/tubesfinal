// package main

// import "database/sql"

// type DeliveryRepository struct {
// 	db *sql.DB
// }

// func NewDeliveryRepository(
// 	db *sql.DB,
// ) *DeliveryRepository {

// 	return &DeliveryRepository{
// 		db: db,
// 	}
// }

// func (r *DeliveryRepository) Create(
// 	d *Delivery,
// ) error {

// 	_, err := r.db.Exec(
// 		`INSERT INTO deliveries
// 		(resi,courier_id,assigned_zone,status)
// 		VALUES(?,?,?,?)`,
// 		d.Resi,
// 		d.CourierID,
// 		d.AssignedZone,
// 		d.Status,
// 	)

// 	return err
// }

// func (r *DeliveryRepository) GetAll() ([]Delivery, error) {

// 	rows, err := r.db.Query(`
// 		SELECT
// 			resi,
// 			courier_id,
// 			assigned_zone,
// 			status,
// 			created_at
// 		FROM deliveries
// 		ORDER BY id DESC
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var deliveries []Delivery

// 	for rows.Next() {

// 		var d Delivery

// 		err := rows.Scan(
// 			&d.Resi,
// 			&d.CourierID,
// 			&d.AssignedZone,
// 			&d.Status,
// 			&d.CreatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		deliveries = append(deliveries, d)
// 	}

// 	return deliveries, nil
// }

// func (r *DeliveryRepository) GetByResi(resi string) (*Delivery, error) {

// 	var d Delivery

// 	err := r.db.QueryRow(`
// 		SELECT
// 			resi,
// 			courier_id,
// 			assigned_zone,
// 			status,
// 			created_at
// 		FROM deliveries
// 		WHERE resi = ?
// 	`, resi).Scan(
// 		&d.Resi,
// 		&d.CourierID,
// 		&d.AssignedZone,
// 		&d.Status,
// 		&d.CreatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &d, nil
// }

package main

import "database/sql"

type DeliveryRepository struct {
	db *sql.DB
}

func NewDeliveryRepository(db *sql.DB) *DeliveryRepository {
	return &DeliveryRepository{
		db: db,
	}
}

func (r *DeliveryRepository) Create(d *Delivery) error {

	_, err := r.db.Exec(
		`
		INSERT INTO deliveries
		(
			resi,
			courier_id,
			nama_penerima,
			no_telp_penerima,
			alamat_penerima,
			berat,
			assigned_zone,
			status
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`,
		d.Resi,
		d.CourierID,
		d.NamaPenerima,
		d.NoTelpPenerima,
		d.AlamatPenerima,
		d.Berat,
		d.AssignedZone,
		d.Status,
	)

	return err
}

func (r *DeliveryRepository) GetAll() ([]Delivery, error) {

	rows, err := r.db.Query(`
		SELECT
			courier_id,
			resi,
			nama_penerima,
			no_telp_penerima,
			alamat_penerima,
			berat,
			status,
			assigned_zone,
			created_at,
			delivered_at
		FROM deliveries
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deliveries []Delivery

	for rows.Next() {

		var d Delivery

		err := rows.Scan(
			&d.CourierID,
			&d.Resi,
			&d.NamaPenerima,
			&d.NoTelpPenerima,
			&d.AlamatPenerima,
			&d.Berat,
			&d.Status,
			&d.AssignedZone,
			&d.CreatedAt,
			&d.DeliveredAt,
		)

		if err != nil {
			return nil, err
		}

		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

func (r *DeliveryRepository) GetByResi(resi string) (*Delivery, error) {

	var d Delivery

	err := r.db.QueryRow(`
		SELECT
			courier_id,
			resi,
			nama_penerima,
			no_telp_penerima,
			alamat_penerima,
			berat,
			status,
			assigned_zone,
			created_at,
			delivered_at
		FROM deliveries
		WHERE resi = ?
	`, resi).Scan(
		&d.CourierID,
		&d.Resi,
		&d.NamaPenerima,
		&d.NoTelpPenerima,
		&d.AlamatPenerima,
		&d.Berat,
		&d.Status,
		&d.AssignedZone,
		&d.CreatedAt,
		&d.DeliveredAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

// 🔥 TAMBAHKAN: Update status delivery
func (r *DeliveryRepository) UpdateStatus(resi string, status string) error {
	_, err := r.db.Exec(
		`UPDATE deliveries SET status = ? WHERE resi = ?`,
		status,
		resi,
	)
	return err
}
