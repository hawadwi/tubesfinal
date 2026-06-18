package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Struktur bantuan untuk membaca JSON kiriman dari order-service
type OrderPayload struct {
	OrderID int    `json:"order_id"`
	UserID  int    `json:"user_id"`
	Resi    string `json:"resi"`
	Berat   int    `json:"berat"`
}

func StartOrderConsumer() {
	msgs, err := rabbitChannel.Consume(
		"payment_queue", // Nama queue yang didengar
		"",              // consumer tags
		true,            // auto-ack (pesan langsung dianggap selesai setelah dibaca)
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Printf("Gagal memulai consumer payment: %s", err)
		return
	}

	// Jalankan di dalam goroutine (background process) supaya tidak memblock main()
	go func() {
		log.Println("[*] Payment Service siap menerima data order...")
		for d := range msgs {
			var order OrderPayload
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("Gagal decode data order di payment: %s", err)
				continue
			}

			log.Printf("[RabbitMQ Payment] Menerima data order baru! ID: %d, Resi: %s", order.OrderID, order.Resi)

			// Simulasi kalkulasi harga sederhana (misal 1kg = Rp 10.000)
			amount := order.Berat * 10000
			if amount <= 0 {
				amount = 15000 // default minimum payment
			}

			// Generate ID transaksi unik
			txID := fmt.Sprintf("TX-%s", order.Resi)

			// Panggil repository untuk menyimpan data ke tabel transactions secara otomatis
			err = paymentRepo.SaveTransaction(txID, int(order.OrderID), amount)
			if err != nil {
				log.Printf("Gagal menyimpan transaksi otomatis ke DB: %s", err)
			} else {
				log.Printf("[DB Payment] Sukses membuat invoice otomatis %s senilai Rp %d dengan status 'pending'", txID, amount)
			}
		}
	}()
}

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func ConnectRabbitMQ() {

	var err error

	// Retry koneksi RabbitMQ
	for i := 0; i < 10; i++ {

		log.Printf("Connecting RabbitMQ... (%d/10)", i+1)

		rabbitConn, err = amqp.Dial(
			"amqp://guest:guest@rabbitmq:5672/",
		)

		if err == nil {
			break
		}

		log.Println("RabbitMQ belum siap:", err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Gagal konek RabbitMQ:", err)
	}

	rabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// Membuat queue jika belum ada
	_, err = rabbitChannel.QueueDeclare(
		"payment_queue",
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("RabbitMQ Connected")
}
