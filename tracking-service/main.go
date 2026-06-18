package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB
	var err error

	// ==========================
	// Ambil konfigurasi dari ENV
	// ==========================
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "mysql"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "root"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "trackingdb"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbname,
	)

	fmt.Println("Using DSN:", dsn)

	// ==========================
	// Retry koneksi MySQL
	// ==========================
	for i := 0; i < 10; i++ {

		db, err = sql.Open("mysql", dsn)

		if err == nil {

			err = db.Ping()

			if err == nil {
				break
			}
		}

		fmt.Println("Waiting MySQL...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	ConnectRabbitMQ()

	StartConsumer()

	fmt.Println("RabbitMQ Connected")

	// ==========================
	// CREATE TABLE OTOMATIS
	// ==========================
	query := `
	CREATE TABLE IF NOT EXISTS tracking_events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		resi VARCHAR(100),
		lokasi VARCHAR(255),
		event VARCHAR(255),
		timestamp DATETIME
	)
	`

	_, err = db.Exec(query)

	if err != nil {
		panic(err)
	}

	fmt.Println("TRACKING TABLE READY")

	trackingRepo = MySQLRepository{
		DB: db,
	}

	// Endpoint Tracking
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy"}`))
	})
	
	http.HandleFunc("/tracking", getTrackingHandler)
	http.HandleFunc("/trackings", getAllTrackingsHandler)
	http.HandleFunc("/tracking/event", insertTrackingEventHandler)

	// Endpoint Map
	http.HandleFunc("/distance", calculateDistanceHandler)
	http.HandleFunc("/route", calculateRouteHandler)
	http.HandleFunc("/location", getCourierLocationHandler)

	fmt.Println("Tracking Service running on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
