// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func ConnectDB() {

// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")

// 	// default untuk local / test
// 	if host == "" {
// 		host = "localhost"
// 	}

// 	if port == "" {
// 		port = "3306"
// 	}

// 	dsn := fmt.Sprintf(
// 		"root:root@tcp(%s:%s)/tubesdb",
// 		host,
// 		port,
// 	)

// 	var err error

// 	for i := 0; i < 10; i++ {

// 		DB, err = sql.Open("mysql", dsn)

// 		if err == nil {
// 			err = DB.Ping()
// 		}

// 		if err == nil {
// 			fmt.Println("REPORT DB CONNECTED")
// 			return
// 		}

// 		fmt.Println("Waiting MySQL...", err)
// 		time.Sleep(5 * time.Second)
// 	}
// }

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
	// 🔥 DIBUAT DINAMIS: Mengambil semua konfigurasi dari Environment Variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Fallback untuk local development / testing jika env kosong
	if host == "" {
		host = "mysql" // Diarahkan ke container mysql secara default
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
		dbname = "reportdb" // Sesuai dengan konfigurasi docker-compose
	}

	// Gunakan parseTime=true agar tipe data DATE/DATETIME MySQL aman saat di-scan ke time.Time Go
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbname,
	)

	var err error

	for i := 0; i < 15; i++ { // Naikkan ke 15 kali retry agar lebih aman menunggu MySQL up
		DB, err = sql.Open("mysql", dsn)

		if err == nil {
			err = DB.Ping()
		}

		if err == nil {
			fmt.Println("REPORT DB CONNECTED SUCCESSFULLY!")
			return
		}

		fmt.Printf("Waiting MySQL at %s:%s... Error: %v\n", host, port, err)
		time.Sleep(3 * time.Second)
	}

	panic(fmt.Sprintf("Failed to connect to database after several retries: %v", err))
}
