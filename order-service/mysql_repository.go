// package main

// type MySQLRepository struct{}

// func (r MySQLRepository) Save(order Order) error {

// 	query := `
// 	INSERT INTO orders
// 	(user_id, nama_barang, berat, dimensi,
// 	jenis, alamat_pengirim, alamat_penerima, status)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	_, err := DB.Exec(
// 		query,
// 		order.UserID,
// 		order.NamaBarang,
// 		order.Berat,
// 		order.Dimensi,
// 		order.Jenis,
// 		order.AlamatPengirim,
// 		order.AlamatPenerima,
// 		order.Status,
// 	)

// 	return err
// }

// package main

// type MySQLRepository struct{}

// func (r MySQLRepository) Save(order Order) error {

// 	query := `
// 	INSERT INTO orders
// 	(
// 		user_id,
// 		resi,
// 		nama_barang,
// 		berat,
// 		dimensi,
// 		jenis,
// 		alamat_pengirim,
// 		alamat_penerima,
// 		status,
// 		eta
// 	)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	_, err := DB.Exec(
// 		query,
// 		order.UserID,
// 		order.Resi,
// 		order.NamaBarang,
// 		order.Berat,
// 		order.Dimensi,
// 		order.Jenis,
// 		order.AlamatPengirim,
// 		order.AlamatPenerima,
// 		order.Status,
// 		order.ETA,
// 	)

// 	return err
// }

// package main

// type MySQLRepository struct{}

// func (r MySQLRepository) GetByID(id int) (*Order, error) {

// 	query := `
// 	SELECT
// 		order_id,
// 		user_id,
// 		resi,
// 		nama_barang,
// 		berat,
// 		dimensi,
// 		jenis,
// 		alamat_pengirim,
// 		alamat_penerima,
// 		status,
// 		eta
// 	FROM orders
// 	WHERE order_id = ?
// 	`

// 	var o Order

// 	err := DB.QueryRow(query, id).Scan(
// 		&o.OrderID,
// 		&o.UserID,
// 		&o.Resi,
// 		&o.NamaBarang,
// 		&o.Berat,
// 		&o.Dimensi,
// 		&o.Jenis,
// 		&o.AlamatPengirim,
// 		&o.AlamatPenerima,
// 		&o.Status,
// 		&o.ETA,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &o, nil
// }

// func (r MySQLRepository) UpdateStatus(
// 	id int,
// 	status string,
// ) error {

// 	query := `
// 	UPDATE orders
// 	SET status = ?
// 	WHERE order_id = ?
// 	`

// 	_, err := DB.Exec(
// 		query,
// 		status,
// 		id,
// 	)

// 	return err
// }

// // 1. UBAH parameter dari (order Order) menjadi (order *Order) agar menggunakan pointer
// func (r MySQLRepository) Save(order *Order) error {

// 	query := `
// 	INSERT INTO orders
// 	(
// 		user_id,
// 		resi,
// 		nama_barang,
// 		berat,
// 		dimensi,
// 		jenis,
// 		alamat_pengirim,
// 		alamat_penerima,
// 		status,
// 		eta
// 	)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	// 2. Tangkap hasil result dari DB.Exec
// 	result, err := DB.Exec(
// 		query,
// 		order.UserID,
// 		order.Resi,
// 		order.NamaBarang,
// 		order.Berat,
// 		order.Dimensi,
// 		order.Jenis,
// 		order.AlamatPengirim,
// 		order.AlamatPenerima,
// 		order.Status,
// 		order.ETA,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	// 3. 🔥 AMBIL ID TERAKHIR YANG DI-GENERATE MYSQL DAN MASUKKAN KE STRUCT
// 	lastID, err := result.LastInsertId()
// 	if err == nil {
// 		order.OrderID = int(lastID) // Nilai order_id sekarang resmi terisi!
// 	}

// 	return nil
// }

// // Tambahkan fungsi ini di paling bawah mysql_repository.go
// func (r MySQLRepository) GetAll() ([]Order, error) {
// 	query := `
// 	SELECT
// 		order_id,
// 		user_id,
// 		resi,
// 		nama_barang,
// 		berat,
// 		dimensi,
// 		jenis,
// 		alamat_pengirim,
// 		alamat_penerima,
// 		status,
// 		eta
// 	FROM orders
// 	`

// 	rows, err := DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var orders []Order
// 	for rows.Next() {
// 		var o Order
// 		err := rows.Scan(
// 			&o.OrderID,
// 			&o.UserID,
// 			&o.Resi,
// 			&o.NamaBarang,
// 			&o.Berat,
// 			&o.Dimensi,
// 			&o.Jenis,
// 			&o.AlamatPengirim,
// 			&o.AlamatPenerima,
// 			&o.Status,
// 			&o.ETA,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		orders = append(orders, o)
// 	}

// 	// Jika di database kosong, return slice kosong [] bukan null
// 	if orders == nil {
// 		orders = []Order{}
// 	}

// 	return orders, nil
// }

package main

type MySQLRepository struct{}

func (r MySQLRepository) GetByID(id int) (*Order, error) {

	query := `
	SELECT
		order_id,
		user_id,
		resi,
		nama_barang,
		berat,
		dimensi,
		jenis,
		alamat_pengirim,
		alamat_penerima,
		nama_penerima,      -- 🔥 Tambahkan kolom ke query
		no_telp_penerima,   -- 🔥 Tambahkan kolom ke query
		status,
		eta
	FROM orders
	WHERE order_id = ?
	`

	var o Order

	err := DB.QueryRow(query, id).Scan(
		&o.OrderID,
		&o.UserID,
		&o.Resi,
		&o.NamaBarang,
		&o.Berat,
		&o.Dimensi,
		&o.Jenis,
		&o.AlamatPengirim,
		&o.AlamatPenerima,
		&o.NamaPenerima,   // 🔥 Scan variabel baru
		&o.NoTelpPenerima, // 🔥 Scan variabel baru
		&o.Status,
		&o.ETA,
	)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r MySQLRepository) UpdateStatus(
	id int,
	status string,
) error {

	query := `
	UPDATE orders
	SET status = ?
	WHERE order_id = ?
	`

	_, err := DB.Exec(
		query,
		status,
		id,
	)

	return err
}

func (r MySQLRepository) Save(order *Order) error {

	query := `
	INSERT INTO orders
	(
		user_id,
		resi,
		nama_barang,
		berat,
		dimensi,
		jenis,
		alamat_pengirim,
		alamat_penerima,
		nama_penerima,      -- 🔥 Tambahkan daftar kolom
		no_telp_penerima,   -- 🔥 Tambahkan daftar kolom
		status,
		eta
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) -- 🔥 Jadi 12 placeholder tanda tanya (?)
	`

	result, err := DB.Exec(
		query,
		order.UserID,
		order.Resi,
		order.NamaBarang,
		order.Berat,
		order.Dimensi,
		order.Jenis,
		order.AlamatPengirim,
		order.AlamatPenerima,
		order.NamaPenerima,   // 🔥 Masukkan nilai data penerima
		order.NoTelpPenerima, // 🔥 Masukkan nilai data penerima
		order.Status,
		order.ETA,
	)

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err == nil {
		order.OrderID = int(lastID)
	}

	return nil
}

func (r MySQLRepository) GetAll() ([]Order, error) {
	query := `
	SELECT
		order_id,
		user_id,
		resi,
		nama_barang,
		berat,
		dimensi,
		jenis,
		alamat_pengirim,
		alamat_penerima,
		nama_penerima,      -- 🔥 Tambahkan kolom ke query
		no_telp_penerima,   -- 🔥 Tambahkan kolom ke query
		status,
		eta
	FROM orders
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		err := rows.Scan(
			&o.OrderID,
			&o.UserID,
			&o.Resi,
			&o.NamaBarang,
			&o.Berat,
			&o.Dimensi,
			&o.Jenis,
			&o.AlamatPengirim,
			&o.AlamatPenerima,
			&o.NamaPenerima,   // 🔥 Scan variabel baru
			&o.NoTelpPenerima, // 🔥 Scan variabel baru
			&o.Status,
			&o.ETA,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	if orders == nil {
		orders = []Order{}
	}

	return orders, nil
}
