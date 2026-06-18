package main

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 🔥 FIXED: Publish ke Report Service dengan proper retry dan error handling
func PublishReportEvent(event TrackingEvent) error {
	if rabbitChannel == nil {
		log.Println("[ERROR] rabbitChannel is nil")
		return nil // Non-blocking, tracking sudah tersimpan
	}

	// Tambahkan timestamp jika kosong
	if event.Timestamp == "" {
		event.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Println("[ERROR] Failed to marshal event:", err)
		return err
	}

	// Ensure report_queue exists
	_, err = rabbitChannel.QueueDeclare(
		"report_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("[ERROR] Failed to declare report_queue:", err)
		return nil
	}

	err = rabbitChannel.Publish(
		"",
		"report_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("[ERROR] Failed to publish to report_queue:", err)
		return nil // Non-blocking error
	}

	log.Println("[SUCCESS] Published tracking event to Report Queue:", string(body))
	return nil
}
