package main

import (
	"testing"
)

func TestStartSortingSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        123,
		Resi:          "RES001",
		NamaBarang:    "Laptop",
		Berat:         2,
		WarehouseZone: "Jakarta",
		Status:        "pending",
	}

	err := service.StartSorting(pkg)
	if err != nil {
		t.Errorf("StartSorting failed: %v", err)
	}

	if pkg.Status != "sorting" {
		t.Errorf("Expected status 'sorting', got '%s'", pkg.Status)
	}
}

func TestStartSortingNil(t *testing.T) {
	service := NewSortingService()

	err := service.StartSorting(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestStartSortingEmptyResi(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        123,
		Resi:          "",
		WarehouseZone: "Jakarta",
		Status:        "pending",
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestStartSortingEmptyWarehouseZone(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        123,
		Resi:          "RES001",
		WarehouseZone: "",
		Status:        "pending",
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for empty warehouse_zone, got nil")
	}
}

func TestStartSortingInvalidStatus(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        123,
		Resi:          "RES001",
		WarehouseZone: "Jakarta",
		Status:        "ready",
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for non-pending package, got nil")
	}
}

func TestCompleteSortingSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		Status:        "sorting",
		WarehouseZone: "Jakarta",
	}

	err := service.CompleteSorting(pkg)
	if err != nil {
		t.Errorf("CompleteSorting failed: %v", err)
	}

	if pkg.Status != "ready" {
		t.Errorf("Expected status 'ready', got '%s'", pkg.Status)
	}

	if pkg.SortedAt == nil {
		t.Error("SortedAt should not be nil")
	}
}

func TestCompleteSortingNil(t *testing.T) {
	service := NewSortingService()

	err := service.CompleteSorting(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestCompleteSortingNotSorting(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		Status:        "pending",
		WarehouseZone: "Jakarta",
	}

	err := service.CompleteSorting(pkg)
	if err == nil {
		t.Error("Expected error for non-sorting package, got nil")
	}
}

/* func TestGetPendingPackagesSuccess(t *testing.T) {
	service := NewSortingService()

	packages := []Package{
		{Resi: "RES001", Status: "pending"},
		{Resi: "RES002", Status: "sorting"},
		{Resi: "RES003", Status: "pending"},
	}

	pending := service.GetPendingPackages(packages)
	if len(pending) != 2 {
		t.Errorf("Expected 2 pending packages, got %d", len(pending))
	}
}

func TestGetPendingPackagesEmpty(t *testing.T) {
	service := NewSortingService()

	packages := []Package{
		{Resi: "RES001", Status: "sorting"},
		{Resi: "RES002", Status: "ready"},
	}

	pending := service.GetPendingPackages(packages)
	if len(pending) != 0 {
		t.Errorf("Expected 0 pending packages, got %d", len(pending))
	}
}

func TestValidatePackageSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		UserID:        123,
		Berat:         2,
		WarehouseZone: "Jakarta",
	}

	err := service.ValidatePackage(pkg)
	if err != nil {
		t.Errorf("ValidatePackage failed: %v", err)
	}
}

func TestValidatePackageNil(t *testing.T) {
	service := NewSortingService()

	err := service.ValidatePackage(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestValidatePackageEmptyResi(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "",
		UserID:        123,
		Berat:         2,
		WarehouseZone: "Jakarta",
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestValidatePackageInvalidUserID(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		UserID:        0,
		Berat:         2,
		WarehouseZone: "Jakarta",
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for invalid user_id, got nil")
	}
}

func TestValidatePackageInvalidWeight(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		UserID:        123,
		Berat:         0,
		WarehouseZone: "Jakarta",
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for zero weight, got nil")
	}
}

func TestValidatePackageEmptyWarehouseZone(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "RES001",
		UserID:        123,
		Berat:         2,
		WarehouseZone: "",
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for empty warehouse_zone, got nil")
	}
} */

func BenchmarkStartSorting(b *testing.B) {
	service := NewSortingService()
	pkg := &Package{
		UserID:        123,
		Resi:          "RES001",
		WarehouseZone: "Jakarta",
		Status:        "pending",
	}

	for i := 0; i < b.N; i++ {
		service.StartSorting(pkg)
		pkg.Status = "pending"
	}
}

func BenchmarkValidatePackage(b *testing.B) {
	service := NewSortingService()
	pkg := &Package{
		Resi:          "RES001",
		UserID:        123,
		Berat:         2,
		WarehouseZone: "Jakarta",
	}

	for i := 0; i < b.N; i++ {
		service.ValidatePackage(pkg)
	}
}
