package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Service Running"))
}

func main() {
	ConnectDB()

	// 1. Daftarkan semua route SEPERTI BIASA (tanpa dibungkus enableCORS)
	http.HandleFunc("/", home)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/check-admin", checkAdminRoleHandler) // 🔥 TAMBAHKAN INI

	// 2. Buat default ServeMux untuk menampung semua handler di atas
	mux := http.DefaultServeMux

	// 3. Bungkus seluruh mux dengan fungsi CORS secara global
	// Karena enableCORS kamu menerima http.HandlerFunc, kita adaptasi sedikit:
	globalCORS := enableCORS(mux.ServeHTTP)

	fmt.Println("User Service running on :8081")
	// 4. Masukkan globalCORS ke ListenAndServe, BUKAN nil
	http.ListenAndServe(":8081", globalCORS)
}
