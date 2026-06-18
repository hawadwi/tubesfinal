package mq

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect() (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			return conn, nil
		}

		log.Printf("RabbitMQ not ready (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}