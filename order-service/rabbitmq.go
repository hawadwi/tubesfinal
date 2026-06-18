// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"time"

// 	amqp "github.com/rabbitmq/amqp091-go"
// )

// // Fungsi untuk mengirim data order ke RabbitMQ
// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Printf("[RabbitMQ Error] %s: %s", msg, err)
// 	}
// }

// func PublishOrderToRabbitMQ(orderData interface{}) {
// 	// Koneksi ke RabbitMQ.
// 	// Saat development lokal (docker-compose), gunakan "amqp://guest:guest@localhost:5672/"
// 	// Catatan untuk temanmu yang handle K8s/Docker: jika nama container rabbitmq di docker-compose adalah "rabbitmq",
// 	// ganti "localhost" menjadi "rabbitmq" saat dideploy.
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		failOnError(err, "Gagal koneksi ke RabbitMQ")
// 		return
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		failOnError(err, "Gagal membuka channel")
// 		return
// 	}
// 	defer ch.Close()

// 	// Daftarkan antrean (queue) tempat gudang-service akan mendengarkan
// 	q, err := ch.QueueDeclare(
// 		"order_queue", // Nama antrean
// 		true,          // durable (antrean tetap ada meski rabbitmq restart)
// 		false,         // delete when unused
// 		false,         // exclusive
// 		false,         // no-wait
// 		nil,           // arguments
// 	)
// 	if err != nil {
// 		failOnError(err, "Gagal mendeklarasikan queue")
// 		return
// 	}

// 	// Ubah data object order menjadi JSON string / bytes
// 	body, err := json.Marshal(orderData)
// 	if err != nil {
// 		failOnError(err, "Gagal mengubah object order menjadi JSON")
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Kirim pesan ke antrean
// 	err = ch.PublishWithContext(ctx,
// 		"",     // exchange
// 		q.Name, // routing key (nama queue)
// 		false,  // mandatory
// 		false,  // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		})

// 	if err != nil {
// 		failOnError(err, "Gagal mempublikasikan pesan ke RabbitMQ")
// 		return
// 	}

// 	log.Printf("[RabbitMQ] Berhasil mengirim data order ke gudang-service: %s", body)
// }

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Fungsi untuk menangani error logs dari RabbitMQ
func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("[RabbitMQ Error] %s: %s", msg, err)
	}
}

func PublishOrderToRabbitMQ(orderData interface{}) {
	// 1. Ambil URL RabbitMQ dari environment variable Docker Compose
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	// 2. Jika kosong (misal running manual tanpa docker), fallback ke amqp di localhost
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@rabbitmq:5672/" // Mengarah ke nama container 'rabbitmq' di Docker
	}

	// Koneksi ke RabbitMQ menggunakan URL dinamis
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		failOnError(err, "Gagal koneksi ke RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Gagal membuka channel")
		return
	}
	defer ch.Close()

	// Daftarkan antrean (queue) tempat gudang-service akan mendengarkan
	q, err := ch.QueueDeclare(
		"order_queue", // Nama antrean
		true,          // durable (antrean tetap ada meski rabbitmq restart)
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		failOnError(err, "Gagal mendeklarasikan queue")
		return
	}

	// Ubah data object order menjadi JSON string / bytes
	body, err := json.Marshal(orderData)
	if err != nil {
		failOnError(err, "Gagal mengubah object order menjadi JSON")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Kirim pesan ke antrean
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key (nama queue)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		failOnError(err, "Gagal mempublikasikan pesan ke RabbitMQ")
		return
	}

	log.Printf("[RabbitMQ] Berhasil mengirim data order ke gudang-service: %s", body)
}

// Tambahkan fungsi ini di paling bawah order-service/rabbitmq.go

// Tambahkan fungsi ini di paling bawah order-service/rabbitmq.go
func PublishToPaymentRabbitMQ(orderData interface{}) {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		failOnError(err, "Gagal koneksi ke RabbitMQ untuk Payment")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Gagal membuka channel untuk Payment")
		return
	}
	defer ch.Close()

	// Daftarkan antrean payment_queue
	q, err := ch.QueueDeclare(
		"payment_queue", // Menembak ke queue milik payment-service
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Gagal mendeklarasikan payment_queue")
		return
	}

	body, err := json.Marshal(orderData)
	if err != nil {
		failOnError(err, "Gagal mengubah object order menjadi JSON untuk Payment")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		failOnError(err, "Gagal mempublikasikan pesan ke payment_queue")
		return
	}

	log.Printf("[RabbitMQ] Berhasil mengirim data order ke payment-service: %s", body)
}
