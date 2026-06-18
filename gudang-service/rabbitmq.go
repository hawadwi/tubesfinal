package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("[RabbitMQ Error] %s: %s", msg, err)
	}
}

func PublishToTrackingQueue(resi string, event string) {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		failOnError(err, "Gagal koneksi ke RabbitMQ untuk Tracking")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Gagal membuka channel untuk Tracking")
		return
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"tracking_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Gagal mendeklarasikan tracking_queue")
		return
	}

	eventData := map[string]interface{}{
		"resi":      resi,
		"lokasi":    "Gudang Service",
		"event":     event,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	body, _ := json.Marshal(eventData)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		"tracking_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		failOnError(err, "Gagal publish ke tracking_queue")
		return
	}

	log.Printf("[RabbitMQ] Event tracking published: %s - %s", resi, event)
}