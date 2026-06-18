package main

type Order struct {
	OrderID        int    `json:"order_id"`
	UserID         int    `json:"user_id"`
	Resi           string `json:"resi"`
	NamaBarang     string `json:"nama_barang"`
	Berat          int    `json:"berat"`
	Dimensi        string `json:"dimensi"`
	Jenis          string `json:"jenis"`
	AlamatPengirim string `json:"alamat_pengirim"`
	AlamatPenerima string `json:"alamat_penerima"`
	NamaPenerima   string `json:"nama_penerima"`    // 🔥 TAMBAHKAN INI
	NoTelpPenerima string `json:"no_telp_penerima"` // 🔥 TAMBAHKAN INI
	Status         string `json:"status"`
	ETA            string `json:"eta"`
}
