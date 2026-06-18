package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	mq "github.com/hawadwi/courier-service/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var db *sql.DB

func InitDB() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "mysql"
	}
	if port == "" {
		port = "3306"
	}
	if user == "" {
		user = "root"
	}
	if password == "" {
		password = "root"
	}
	if dbname == "" {
		log.Fatal("DB_NAME required")
	}

	connStr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"

	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}

		log.Printf("DB not ready, retrying... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Courier database connected")
	return nil
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// ======================
// RABBIT CONNECTION
// ======================
func connectRabbitMQ() *amqp.Connection {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@rabbitmq:5672/"
	}

	var conn *amqp.Connection
	var err error

	for i := 0; i < 15; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			return conn
		}

		log.Printf("RabbitMQ not ready, retrying... (%d/15)", i+1)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("failed to connect rabbitmq:", err)
	return nil
}

func main() {

	// ======================
	// DB INIT
	// ======================
	if err := InitDB(); err != nil {
		log.Fatalf("database failed: %v", err)
	}
	defer CloseDB()

	// ======================
	// SERVICE INIT
	// ======================
	repo := NewDeliveryRepository(GetDB())
	service := NewCourierService(repo)

	// ======================
	// RABBITMQ INIT
	// ======================
	conn := connectRabbitMQ()
	defer conn.Close()

	// ======================
	// START CONSUMER
	// ======================
	go func() {
		time.Sleep(3 * time.Second)

		err := mq.StartConsumer(conn, func(event mq.DeliveryEvent) error {
			return service.ProcessDelivery(event)
		})

		if err != nil {
			log.Println("consumer error:", err)
		}
	}()

	// ======================
	// HTTP SERVER
	// ======================
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to create channel:", err)
	}

	handler := NewCourierHandler(service, repo, ch)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	
	http.HandleFunc("/deliveries", GetDeliveries(repo))
	http.HandleFunc("/delivery/complete", handler.CompleteDelivery)
	http.HandleFunc("/delivery", handler.StartDelivery)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8086"
	}

	log.Printf("Courier Service running on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
