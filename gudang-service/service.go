package main

import (
	"encoding/json"
	"errors"
	"time"
)

type SortingService struct {
	repo *PackageRepository
}

func NewSortingService(repo *PackageRepository) *SortingService {
	return &SortingService{repo: repo}
}

// START SORT
func (s *SortingService) StartSorting(pkg *Package) error {
	if pkg.Resi == "" {
		return errors.New("resi kosong")
	}
	pkg.Status = "sorting"
	return nil
}

// COMPLETE SORT
func (s *SortingService) CompleteSorting(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	now := time.Now()

	pkg.Status = "ready"
	pkg.SortedAt = &now

	event := map[string]any{
		"event": "package.ready",
		"data": map[string]any{
			"resi":           pkg.Resi,
			"nama_penerima":     pkg.NamaPenerima,        // NEW
			"no_telp_penerima":  pkg.NoTelpPenerima,      // NEW
			"alamat_penerima":   pkg.AlamatPenerima,
			"berat":             pkg.Berat,               // NEW (was missing)
			"warehouse_zone": pkg.WarehouseZone,
			"status":         "ready",
			"sorted_at":      now,
		},
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.repo.SaveOutbox("package.ready", string(b))
}