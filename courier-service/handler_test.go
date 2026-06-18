package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test start delivery

func TestStartDelivery_Success(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	body := []byte(`{
	    "resi": "RESI001",
	    "courier_id": 1,
	    "assigned_zone": "Jakarta"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()

	// expected response ketika request berhasil
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestStartDelivery_InvalidJSON(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	// isi dengan format JSON yang tidak valid
	body := []byte(``)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()

	// expected response ketika format JSON salah
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

func TestStartDelivery_MissingField(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	body := []byte(`{
		"resi": "", // kosongkan nomor resi
		"courier_id": 0, // isi dengan courier_id tidak valid/kosong
		"assigned_zone": "" // kosongkan zona pengiriman
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()

	// expected response ketika ada field yang kosong
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

// test get delivery

func TestGetCourierDeliveries_Success(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	// isi query parameter dengan courier_id valid
	req := httptest.NewRequest( http.MethodGet, "/courier/deliveries?courier_id=1", nil, )
	
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()

	// expected response ketika data berhasil diambil
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestGetCourierDeliveries_InvalidID(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	// isi courier_id dengan format yang tidak valid
	req := httptest.NewRequest( http.MethodGet, "/courier/deliveries?courier_id=abc", nil, )

	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()

	// expected response ketika courier_id invalid
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

func TestGetCourierDeliveries_MissingID(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	// request tanpa query parameter courier_id
	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries", nil)

	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()

	// expected response ketika courier_id tidak dikirim
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

// test health 

func TestHealth(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	// endpoint health check service
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	w := httptest.NewRecorder()

	handler.Health(w, req)

	res := w.Result()

	// expected response ketika service berjalan normal
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}
