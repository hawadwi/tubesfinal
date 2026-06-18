package main

import "time"

type Delivery struct {
	CourierID      int        `json:"courier_id"`
	Resi           string     `json:"resi"`
	NamaPenerima   string     `json:"nama_penerima"`
	NoTelpPenerima string     `json:"no_telp_penerima"`
	AlamatPenerima string     `json:"alamat_penerima"`
	Berat          int        `json:"berat"`
	Status         string     `json:"status"` // pending, in_delivery, delivered
	AssignedZone   string     `json:"assigned_zone"`
	CreatedAt      time.Time  `json:"created_at"`
	DeliveredAt    *time.Time `json:"delivered_at"`
}

type DeliveryRequest struct {
	Resi           string `json:"resi" binding:"required"`
	CourierID      int    `json:"courier_id" binding:"required"`
	AssignedZone   string `json:"assigned_zone" binding:"required"`
	NamaPenerima   string `json:"nama_penerima"`    // <-- Tambahkan ini
	NoTelpPenerima string `json:"no_telp_penerima"` // <-- Tambahkan ini
	AlamatPenerima string `json:"alamat_penerima"`  // <-- Tambahkan ini
	Berat          int    `json:"berat"`            // <-- Tambahkan ini
}
