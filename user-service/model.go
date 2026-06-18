package main

type User struct {
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"` // Tag "-" memberitahu Go untuk mengabaikan field ini saat konversi ke JSON
	Alamat     string `json:"alamat"`
	Preferensi string `json:"preferensi"`
	Role       string `json:"role"`
}
