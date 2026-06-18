package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func ConnectRabbitMQ() {

	var err error

	log.Println("Connecting RabbitMQ...")

	rabbitConn, err = amqp.Dial(
		"amqp://guest:guest@rabbitmq:5672/",
	)

	if err != nil {
		log.Fatal("Dial error:", err)
	}

	log.Println("RabbitMQ Dial Success")

	rabbitChannel, err = rabbitConn.Channel()

	if err != nil {
		log.Fatal("Channel error:", err)
	}

	log.Println("Channel Success")

	// Queue untuk menerima event dari Order/Gudang/Courier
	_, err = rabbitChannel.QueueDeclare(
		"tracking_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("QueueDeclare error:", err)
	}

	// Queue untuk mengirim event ke Report Service
	_, err = rabbitChannel.QueueDeclare(
		"report_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("QueueDeclare report_queue error:", err)
	}

	log.Println("Queue tracking_queue created")
	log.Println("Queue report_queue created")
}