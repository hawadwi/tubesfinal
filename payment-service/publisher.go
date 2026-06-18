package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishPaymentEvent(event PaymentEvent) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = rabbitChannel.Publish(
		"",
		"payment_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	log.Println("Berhasil publish ke RabbitMQ:", string(body))

	return nil
}