package mq

import (
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex
)

func InitRabbitMQ(url string) error {

	var err error

	for i := 0; i < 15; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		log.Printf("RabbitMQ retrying... (%d/15)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("rabbitmq connect failed: %w", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("channel failed: %w", err)
	}

	log.Println("RabbitMQ connected")
	return nil
}

func Publish(queue string, payload string) error {

	mu.Lock()
	defer mu.Unlock()

	if ch == nil {
		return fmt.Errorf("channel not ready")
	}

	_, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
			Timestamp:   time.Now(),
		},
	)
}

func Close() {
	mu.Lock()
	defer mu.Unlock()

	if ch != nil {
		_ = ch.Close()
	}
	if conn != nil {
		_ = conn.Close()
	}
}