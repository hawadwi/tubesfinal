package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 🔥 IMPROVED: Consumer dengan better error handling dan reconnect logic
func StartConsumer() {

	if rabbitChannel == nil {
		log.Fatal("rabbitChannel masih nil")
	}

	// Declare queue untuk memastikan ada
	_, err := rabbitChannel.QueueDeclare(
		"tracking_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to declare tracking_queue:", err)
	}

	msgs, err := rabbitChannel.Consume(
		"tracking_queue",
		"",
		false, // 🔥 FIXED: Changed to false untuk manual ack
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to start consuming:", err)
	}

	go func() {
		for msg := range msgs {

			fmt.Println("Message diterima dari tracking_queue:")
			fmt.Println(string(msg.Body))

			// Decode JSON
			var event TrackingEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Println("[ERROR] Failed to unmarshal event:", err)
				msg.Nack(false, false) // Reject message
				continue
			}

			// Add timestamp if missing
			if event.Timestamp == "" {
				event.Timestamp = time.Now().Format("2006-01-02 15:04:05")
			}

			// ===========================
			// Simpan ke database
			// ===========================

			err = trackingRepo.Insert(event)

			if err != nil {
				log.Println("[ERROR] Gagal simpan ke MySQL:", err)
				msg.Nack(false, true) // Requeue message
				continue
			}

			log.Println("[SUCCESS] Tracking berhasil disimpan:", event.Resi)

			// ===========================
			// Publish ke Report Service
			// ===========================

			err = PublishReportEvent(event)

			if err != nil {
				log.Println("[WARNING] Gagal publish ke report:", err)
				// Continue anyway - tracking sudah tersimpan
			} else {
				log.Println("[SUCCESS] Berhasil publish ke report_queue")
			}

			// Acknowledge message
			msg.Ack(false)

		}

	}()

	log.Println("[INFO] Tracking Consumer Started - Listening to tracking_queue")

}
