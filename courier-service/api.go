package main

import (
	"encoding/json"
	"net/http"
)

func GetDeliveries(repo *DeliveryRepository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resi := r.URL.Query().Get("resi")

		w.Header().Set("Content-Type", "application/json")

		// GET /deliveries?resi=RESI001
		if resi != "" {

			delivery, err := repo.GetByResi(resi)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(delivery)
			return
		}

		// GET /deliveries
		data, err := repo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(data)
	}
}