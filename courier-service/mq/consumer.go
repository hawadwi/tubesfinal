package mq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type DeliveryEvent struct {
	Event string `json:"event"`
	Data  struct {
		Resi          string `json:"resi"`
		WarehouseZone string `json:"warehouse_zone"`
		NamaPenerima      string `json:"nama_penerima"`        // NEW
		NoTelpPenerima    string `json:"no_telp_penerima"`     // NEW
		AlamatPenerima    string `json:"alamat_penerima"`      // NEW
		Berat             int    `json:"berat"`                // NEW
	} `json:"data"`
}

func StartConsumer(conn *amqp.Connection, handler func(DeliveryEvent) error) error {

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"package.ready",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("Courier consumer running")

	go func() {
		for msg := range msgs {

			var event DeliveryEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Println("invalid event:", err)
				continue
			}

			log.Println("RECEIVED EVENT:", event.Data.Resi)

			_ = handler(event)
		}
	}()

	return nil
}
