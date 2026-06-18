package mq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 🔥 TAMBAHKAN: Publisher untuk mengirim tracking event ke tracking-service
func PublishTrackingEvent(conn *amqp.Connection, resi string, event string, location string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Deklarasi tracking_queue
	_, err = ch.QueueDeclare(
		"tracking_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Buat event payload
	trackingEvent := map[string]interface{}{
		"resi":      resi,
		"lokasi":    location,
		"event":     event,
		"timestamp": "",
	}

	body, err := json.Marshal(trackingEvent)
	if err != nil {
		return err
	}

	// Publish ke tracking_queue
	err = ch.Publish(
		"",
		"tracking_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Failed to publish tracking event:", err)
		return err
	}

	log.Println("Published tracking event for resi:", resi)
	return nil
}
