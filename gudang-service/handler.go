package main

import (
	"encoding/json"
	"net/http"
)

type SortingHandler struct {
	service    *SortingService
	repository *PackageRepository
}

func NewSortingHandler(service *SortingService, repository *PackageRepository) *SortingHandler {
	return &SortingHandler{
		service:    service,
		repository: repository,
	}
}

// ======================
// START SORT
// ======================
func (h *SortingHandler) StartSort(w http.ResponseWriter, r *http.Request) {

	var req SortRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	pkg := &Package{
		Resi:       req.Resi,
		NamaBarang: req.NamaBarang,
		Berat:      req.Berat,
		Status:     "sorting",
	}

	if err := h.service.StartSorting(pkg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repository.Create(pkg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pkg)
}

// ======================
// COMPLETE SORT
// ======================
func (h *SortingHandler) CompleteSort(w http.ResponseWriter, r *http.Request) {

	var req CompleteSortRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	pkg, err := h.repository.GetByResi(req.Resi)
	if err != nil {
		http.Error(w, "resi not found", http.StatusNotFound)
		return
	}

	if err := h.repository.CompleteSort(req.Resi); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.CompleteSorting(pkg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(pkg)
}

