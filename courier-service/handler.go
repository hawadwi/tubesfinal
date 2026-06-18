// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/streadway/amqp"
// )

// type CourierServiceInterface interface {
// 	StartDelivery(delivery *Delivery) error
// 	CompleteDelivery(delivery *Delivery) error
// 	GetCourierDeliveries(deliveries []Delivery, courierID int) []Delivery
// 	ValidateDelivery(delivery *Delivery) error
// }

// type CourierHandler struct {
// 	service    CourierServiceInterface
// 	repository *DeliveryRepository
// }

// type CompleteDeliveryRequest struct {
// 	Resi string `json:"resi"`
// }

// func NewCourierHandler(
// 	service CourierServiceInterface,
// 	repository *DeliveryRepository,
// ) *CourierHandler {

// 	return &CourierHandler{
// 		service:    service,
// 		repository: repository,
// 	}
// }

// // POST /delivery
// func (h *CourierHandler) StartDelivery(w http.ResponseWriter, r *http.Request) {
// 	var req DeliveryRequest

// 	// decode request body dari JSON
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	log.Println("COURIER RECEIVED DELIVERY:", req.Resi)

// 	// validasi field
// 	if req.Resi == "" || req.CourierID <= 0 || req.AssignedZone == "" {
// 		http.Error(w, "resi, courier_id, assigned_zone are required", http.StatusBadRequest)
// 		return
// 	}

// 	delivery := &Delivery{
// 		Resi:           req.Resi,
// 		CourierID:      req.CourierID,
// 		AssignedZone:   req.AssignedZone,
// 		NamaPenerima:   req.NamaPenerima,   // NEW
// 		NoTelpPenerima: req.NoTelpPenerima, // NEW
// 		AlamatPenerima: req.AlamatPenerima, // NEW
// 		Berat:          req.Berat,          // NEW
// 		Status:         "pending",
// 		CreatedAt:      time.Now(),
// 	}

// 	// panggil service StartDelivery
// 	if err := h.service.StartDelivery(delivery); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err := h.repository.Create(delivery)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			err.Error(),
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}

// 	// 🔥 TAMBAHKAN PUBLISH KE TRACKING DI SINI
// 	go func() {
// 		event := map[string]interface{}{
// 			"resi":      delivery.Resi,
// 			"lokasi":    "Kurir ID " + strconv.Itoa(delivery.CourierID),
// 			"event":     "Paket sedang dibawa oleh kurir menuju lokasi penerima",
// 			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
// 		}
// 		body, _ := json.Marshal(event)
// 		// Panggil channel RabbitMQ milik courier-service untuk kirim ke "tracking_queue"
// 		CourierRabbitChannel.Publish("", "tracking_queue", false, false, amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		})
// 	}()

// 	w.Header().Set("Content-Type", "application/json")

// 	// kirim response delivery dalam format JSON
// 	json.NewEncoder(w).Encode(delivery)
// }

// // GET /courier/deliveries?courier_id=1
// func (h *CourierHandler) GetCourierDeliveries(w http.ResponseWriter, r *http.Request) {
// 	idStr := r.URL.Query().Get("courier_id")

// 	// validasi query parameter courier_id
// 	if idStr == "" {
// 		http.Error(w, "courier_id is required", http.StatusBadRequest)
// 		return
// 	}

// 	courierID, err := strconv.Atoi(idStr)

// 	// validasi apakah courier_id berupa angka valid
// 	if err != nil || courierID <= 0 {
// 		http.Error(w, "invalid courier_id", http.StatusBadRequest)
// 		return
// 	}

// 	// data delivery
// 	all := []Delivery{
// 		{
// 			Resi:      "", // isi dengan nomor resi
// 			CourierID: 0,  // isi dengan courier_id
// 			Status:    "", // isi dengan status delivery
// 		},
// 		{
// 			Resi:      "", // isi dengan nomor resi
// 			CourierID: 0,  // isi dengan courier_id
// 			Status:    "", // isi dengan status delivery
// 		},
// 	}

// 	// ambil data delivery berdasarkan courier_id
// 	result := h.service.GetCourierDeliveries(all, courierID)

// 	w.Header().Set("Content-Type", "application/json")

// 	// kirim response hasil delivery courier
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"courier_id": courierID,
// 		"count":      len(result),
// 		"data":       result,
// 	})
// }

// func (h *CourierHandler) CompleteDelivery(w http.ResponseWriter, r *http.Request) {

// 	var req CompleteDeliveryRequest

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid body", http.StatusBadRequest)
// 		return
// 	}

// 	delivery := &Delivery{
// 		Resi:   req.Resi,
// 		Status: "in_delivery",
// 	}

// 	if err := h.service.CompleteDelivery(delivery); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(delivery)
// }

// // GET /health
// func (h *CourierHandler) Health(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// response health check service
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"status": "healthy", // isi dengan status service
// 	})
// }

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go" // 🔥 FIXED: Disamakan menggunakan library resmi amqp091-go
)

type CourierServiceInterface interface {
	StartDelivery(delivery *Delivery) error
	CompleteDelivery(delivery *Delivery) error
	GetCourierDeliveries(deliveries []Delivery, courierID int) []Delivery
	ValidateDelivery(delivery *Delivery) error
}

type CourierHandler struct {
	service    CourierServiceInterface
	repository *DeliveryRepository
	ch         *amqp091.Channel // 🔥 FIXED: Kita simpan channel RabbitMQ di dalam struct handler
}

type CompleteDeliveryRequest struct {
	Resi string `json:"resi"`
}

// 🔥 FIXED: Tambahkan parameter *amqp091.Channel ke constructor
func NewCourierHandler(
	service CourierServiceInterface,
	repository *DeliveryRepository,
	ch *amqp091.Channel,
) *CourierHandler {

	return &CourierHandler{
		service:    service,
		repository: repository,
		ch:         ch,
	}
}

// POST /delivery
func (h *CourierHandler) StartDelivery(w http.ResponseWriter, r *http.Request) {
	var req DeliveryRequest

	// decode request body dari JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println("COURIER RECEIVED DELIVERY:", req.Resi)

	// validasi field
	if req.Resi == "" || req.CourierID <= 0 || req.AssignedZone == "" {
		http.Error(w, "resi, courier_id, assigned_zone are required", http.StatusBadRequest)
		return
	}

	delivery := &Delivery{
		Resi:           req.Resi,
		CourierID:      req.CourierID,
		AssignedZone:   req.AssignedZone,
		NamaPenerima:   req.NamaPenerima,
		NoTelpPenerima: req.NoTelpPenerima,
		AlamatPenerima: req.AlamatPenerima,
		Berat:          req.Berat,
		Status:         "pending",
		CreatedAt:      time.Now(),
	}

	// panggil service StartDelivery
	if err := h.service.StartDelivery(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repository.Create(delivery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔥 FIXED: Publish menggunakan channel internal struct
	go func() {
		if h.ch == nil {
			log.Println("Warning: RabbitMQ channel is nil, tracking event not sent")
			return
		}

		event := map[string]interface{}{
			"resi":      delivery.Resi,
			"lokasi":    "Kurir ID " + strconv.Itoa(delivery.CourierID),
			"event":     "Paket sedang dibawa oleh kurir menuju lokasi penerima",
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		}
		body, _ := json.Marshal(event)

		// Kirim ke tracking_queue
		err := h.ch.Publish(
			"",               // exchange
			"tracking_queue", // routing key
			false,            // mandatory
			false,            // immediate
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Println("Failed to publish tracking event:", err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(delivery)
}

// GET /courier/deliveries?courier_id=1
func (h *CourierHandler) GetCourierDeliveries(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("courier_id")

	if idStr == "" {
		http.Error(w, "courier_id is required", http.StatusBadRequest)
		return
	}

	courierID, err := strconv.Atoi(idStr)
	if err != nil || courierID <= 0 {
		http.Error(w, "invalid courier_id", http.StatusBadRequest)
		return
	}

	all := []Delivery{
		{Resi: "", CourierID: 0, Status: ""},
		{Resi: "", CourierID: 0, Status: ""},
	}

	result := h.service.GetCourierDeliveries(all, courierID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"courier_id": courierID,
		"count":      len(result),
		"data":       result,
	})
}

func (h *CourierHandler) CompleteDelivery(w http.ResponseWriter, r *http.Request) {
	var req CompleteDeliveryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	// Get current delivery
	delivery, err := h.repository.GetByResi(req.Resi)
	if err != nil || delivery == nil {
		http.Error(w, "delivery not found", http.StatusNotFound)
		return
	}

	// Update status to delivered
	if err := h.repository.UpdateStatus(req.Resi, "delivered"); err != nil {
		http.Error(w, "failed to update delivery", http.StatusInternalServerError)
		return
	}

	// Publish to tracking
	go func() {
		if h.ch == nil {
			log.Println("Warning: RabbitMQ channel is nil")
			return
		}

		event := map[string]interface{}{
			"resi":      req.Resi,
			"lokasi":    "Serah terima ke penerima",
			"event":     "Paket telah berhasil diterima",
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		}
		body, _ := json.Marshal(event)

		err := h.ch.Publish(
			"",
			"tracking_queue",
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Println("Failed to publish tracking event:", err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"resi":   req.Resi,
		"status": "delivered",
		"message": "Delivery completed successfully",
	})
}

// GET /health
func (h *CourierHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}
