package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if host == "" {
		host = "mysql"
	}

	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf(
		"root:root@tcp(%s:%s)/userdb",
		host,
		port,
	)

	var err error

	for i := 0; i < 10; i++ {

		DB, err = sql.Open("mysql", dsn)

		if err == nil {
			err = DB.Ping()
		}

		if err == nil {
			fmt.Println("USER DB CONNECTED")
			createTable()
			return
		}

		fmt.Println("Waiting MySQL...", err)

		time.Sleep(5 * time.Second)
	}

	panic(err)
}

func createTable() {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100),
		password VARCHAR(255),
		role VARCHAR(50),
		alamat TEXT,
		preferensi TEXT
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
