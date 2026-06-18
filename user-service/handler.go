package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("rahasia_super_aman_123")

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	json.NewDecoder(r.Body).Decode(&req)

	u, err := Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"user_id": u.UserID,
		"name":    u.Name,
		"email":   u.Email,
		"role":    u.Role,
	})

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	json.NewDecoder(r.Body).Decode(&req)

	user, err := Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.Email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		w.WriteHeader(401)
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if !VerifyToken(token) {
		w.WriteHeader(403)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if r.Method == "GET" {
		u := GetProfile(id)
		if u == nil {
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(u)
		return
	}

	if r.Method == "PUT" {
		var req User
		json.NewDecoder(r.Body).Decode(&req)

		ok := UpdateProfile(id, req.Alamat, req.Preferensi)
		if !ok {
			w.WriteHeader(404)
			return
		}

		w.Write([]byte(`{"message":"updated"}`))
	}
}

// 🔥 TAMBAHKAN: Admin Handler untuk akses gudang-admin dan kurir-admin
func checkAdminRoleHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if !VerifyToken(token) {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid token"})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	user := GetProfile(userID)
	if user == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	// 🔥 Check if user role is admin
	if user.Role != "admin" && user.Role != "gudang_admin" && user.Role != "kurir_admin" {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{"error": "user is not an admin"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"allowed": true,
		"role":    user.Role,
		"user_id": user.UserID,
	})
}
