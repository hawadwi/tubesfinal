package main

import (
	"errors"
	"time"

	mq "github.com/hawadwi/courier-service/mq"
)

type CourierService struct {
	repository *DeliveryRepository
}

func NewCourierService(repo *DeliveryRepository) *CourierService {
	return &CourierService{repository: repo}
}

func (s *CourierService) StartDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	if d.Resi == "" {
		return errors.New("resi kosong")
	}

	if d.CourierID <= 0 {
		return errors.New("courier invalid")
	}

	if d.Status != "pending" {
		return errors.New("status bukan pending")
	}

	d.Status = "in_delivery"

	return nil
}

func (s *CourierService) CompleteDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	if d.Status != "in_delivery" {
		return errors.New("delivery belum berjalan")
	}

	now := time.Now()

	d.Status = "delivered"
	d.DeliveredAt = &now

	return nil
}

// 🔥 TAMBAHKAN: Method untuk publish ke tracking service
func (s *CourierService) PublishDeliveryUpdate(resi string, status string) {
	// This will be called from handler with global rabbitmq connection
	// Implementation will be in handler layer
}

func (s *CourierService) GetCourierDeliveries(
	deliveries []Delivery,
	courierID int,
) []Delivery {

	var result []Delivery

	for _, d := range deliveries {

		if d.CourierID == courierID {
			result = append(result, d)
		}
	}

	return result
}

func (s *CourierService) ValidateDelivery(delivery *Delivery) error {
	if delivery == nil {
		return errors.New("delivery nil")
	}

	if delivery.Resi == "" {
		return errors.New("resi kosong")
	}

	if delivery.CourierID <= 0 {
		return errors.New("courier_id tidak valid")
	}

	if delivery.AlamatPenerima == "" {
		return errors.New("alamat penerima kosong")
	}

	return nil
}

func (s *CourierService) ProcessDelivery(event mq.DeliveryEvent) error {

	if event.Data.Resi == "" {
		return errors.New("resi kosong")
	}

	// WAJIB: set courier ID (sementara hardcode dulu)
	delivery := &Delivery{
		Resi:           event.Data.Resi,
		CourierID:      1, // <-- FIX PENTING
		AssignedZone:   event.Data.WarehouseZone,
		NamaPenerima:   event.Data.NamaPenerima,   // NEW
		NoTelpPenerima: event.Data.NoTelpPenerima, // NEW
		AlamatPenerima: event.Data.AlamatPenerima, // NEW
		Berat:          event.Data.Berat,          // NEW
		Status:         "in_delivery",
	}

	return s.repository.Create(delivery)
}
