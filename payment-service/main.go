package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB
	var err error

	// Retry koneksi MySQL
	for i := 0; i < 10; i++ {

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
			dbname = "paymentdb"
		}

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			user,
			password,
			host,
			port,
			dbname,
		)

		db, err = sql.Open("mysql", dsn)

		if err == nil {

			err = db.Ping()

			if err == nil {
				break
			}
		}

		fmt.Println("Waiting MySQL...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	ConnectRabbitMQ()
	StartOrderConsumer()

	// ==================================
	// CREATE TABLE OTOMATIS
	// ==================================
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		transaction_id VARCHAR(100) PRIMARY KEY,
		order_id INT,
		amount INT,
		metode VARCHAR(100),
		status VARCHAR(50),
		timestamp DATETIME
	)
	`

	_, err = db.Exec(query)

	if err != nil {
		panic(err)
	}

	fmt.Println("TRANSACTIONS TABLE READY")

	paymentRepo = MySQLRepository{
		DB: db,
	}

	// Endpoint Payment Service
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy"}`))
	})

	http.HandleFunc("/calculate", calculatePaymentHandler)
	http.HandleFunc("/pay", processPaymentHandler)

	fmt.Println("Payment Service running on :8088")
	http.ListenAndServe(":8088", nil)
}
