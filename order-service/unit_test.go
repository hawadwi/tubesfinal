// package main

// type MockValidator struct{}

// func (m MockValidator) CheckUser(userID int, token string) bool {
// 	return true
// }

// type MockOrderRepository struct{}

// func (m MockOrderRepository) Save(order Order) error {
// 	return nil
// }

// func (m MockOrderRepository) GetByID(id int) (*Order, error) {
// 	return &Order{}, nil
// }

// func (m MockOrderRepository) UpdateStatus(
// 	id int,
// 	status string,
// ) error {
// 	return nil
// }

package main

import "testing"

type MockValidator struct{}

func (m MockValidator) CheckUser(userID int, token string) bool {
	return true
}

type MockOrderRepository struct{}

func (m MockOrderRepository) Save(order Order) error {
	return nil
}

func (m MockOrderRepository) GetByID(id int) (*Order, error) {
	return &Order{}, nil
}

func (m MockOrderRepository) UpdateStatus(id int, status string) error {
	return nil
}

func TestCreateOrder(t *testing.T) {

	mockValidator := MockValidator{}
	mockRepo := MockOrderRepository{}

	req := Order{
		UserID:         1,
		NamaBarang:     "Laptop",
		Berat:          2,
		Dimensi:        "10x10",
		Jenis:          "Elektronik",
		AlamatPengirim: "Bandung",
		AlamatPenerima: "Jakarta",
	}

	order, err := CreateOrder(
		req,
		"dummy-token",
		mockValidator,
		mockRepo,
	)

	if err != nil {
		t.Fatal(err)
	}

	if order.Status != "created" {
		t.Fatalf(
			"expected created got %s",
			order.Status,
		)
	}
}
