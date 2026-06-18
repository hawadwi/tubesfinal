package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hawadwi/gudang-service/mq"
)

var db *sql.DB

func InitDB() error {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// default kalau belum pakai env
		dsn = "root:root@tcp(mysql:3306)/gudang?parseTime=true"
	}

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
	}
}

func main() {

	if err := InitDB(); err != nil {
		log.Fatal(err)
	}
	defer CloseDB()

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	if err := mq.InitRabbitMQ(rabbitURL); err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	repo := NewPackageRepository(db)
	service := NewSortingService(repo)
	handler := NewSortingHandler(service, repo)

	go StartOutboxWorker(repo)
	go StartOrderConsumer(repo)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy"}`))
	})

	http.HandleFunc("/sort", handler.StartSort)
	http.HandleFunc("/sort/complete", handler.CompleteSort)
	http.HandleFunc("/packages", GetPackages(repo))

	port := "8085"
	log.Println("running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
