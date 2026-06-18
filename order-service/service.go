// package main

// import "errors"

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// var orders []Order
// var nextID = 1

// var UserServiceURL = "http://host.docker.internal:8081"

// type Validator interface {
// 	CheckUser(userID int, token string) bool
// }

// type RealValidator struct{}

// func (v RealValidator) CheckUser(userID int, token string) bool {
// 	req, _ := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
// 		nil,
// 	)

// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{Timeout: 3 * time.Second}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("PROFILE ERROR:", err)
// 		return false
// 	}

// 	fmt.Println("PROFILE STATUS:", resp.StatusCode)

// 	return resp.StatusCode == 200
// }

// func GenerateResi() string {
// 	return ""
// }

// func CalculateETA() string {
// 	return ""
// }

// func CreateOrder(
// 	req Order,
// 	token string,
// 	validator Validator,
// 	repo OrderRepository,
// ) (Order, error) {

// 	valid := validator.CheckUser(req.UserID, token)

// 	if !valid {
// 		return Order{}, errors.New("user tidak valid")
// 	}

// 	req.Status = "created"

// 	err := repo.Save(req)

// 	if err != nil {
// 		return Order{}, err
// 	}

// 	return req, nil
// }

// func GetOrder(id int) *Order {
// 	return nil
// }

// func UpdateOrderStatus(id int, status string) bool {
// 	return false
// }

// func GetETA(id int) string {
// 	return ""
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// // Docker Compose
// var UserServiceURL = "http://user-service:8081"

// // Local testing
// // var UserServiceURL = "http://localhost:8081"

// type Validator interface {
// 	CheckUser(userID int, token string) bool
// }

// type RealValidator struct{}

// func (v RealValidator) CheckUser(userID int, token string) bool {

// 	req, _ := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
// 		nil,
// 	)

// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{
// 		Timeout: 3 * time.Second,
// 	}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("PROFILE ERROR:", err)
// 		return false
// 	}

// 	fmt.Println("PROFILE STATUS:", resp.StatusCode)

// 	return resp.StatusCode == 200
// }

// func GenerateResi() string {
// 	return fmt.Sprintf(
// 		"LNG-%d",
// 		time.Now().UnixNano(),
// 	)
// }

// func CalculateETA() string {
// 	return "2 days"
// }

// func CreateOrder(
// 	req Order,
// 	token string,
// 	validator Validator,
// 	repo OrderRepository,
// ) (Order, error) {

// 	valid := validator.CheckUser(
// 		req.UserID,
// 		token,
// 	)

// 	if !valid {
// 		return Order{}, errors.New(
// 			"user tidak valid",
// 		)
// 	}

// 	if req.Berat <= 0 {
// 		return Order{}, errors.New(
// 			"berat tidak valid",
// 		)
// 	}

// 	req.Resi = GenerateResi()
// 	req.ETA = CalculateETA()
// 	req.Status = "created"

// 	err := repo.Save(req)

// 	if err != nil {
// 		return Order{}, err
// 	}

// 	return req, nil
// }

// func GetOrder(id int) *Order {
// 	return nil
// }

// func UpdateOrderStatus(
// 	id int,
// 	status string,
// ) bool {
// 	return false
// }

// func GetETA(id int) string {
// 	return ""
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// // Docker Compose
// var UserServiceURL = "http://user-service:8081"

// // Local testing
// // var UserServiceURL = "http://localhost:8081"

// type Validator interface {
// 	CheckUser(userID int, token string) bool
// }

// type RealValidator struct{}

// func (v RealValidator) CheckUser(userID int, token string) bool {

// 	req, _ := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
// 		nil,
// 	)

// 	req.Header.Set("Authorization", "Bearer "+token)

// 	client := &http.Client{
// 		Timeout: 3 * time.Second,
// 	}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("PROFILE ERROR:", err)
// 		return false
// 	}

// 	fmt.Println("PROFILE STATUS:", resp.StatusCode)

// 	return resp.StatusCode == 200
// }

// func GenerateResi() string {
// 	return fmt.Sprintf(
// 		"LNG-%d",
// 		time.Now().UnixNano(),
// 	)
// }

// func CalculateETA() string {
// 	return "2 days"
// }

// func CreateOrder(
// 	req Order,
// 	token string,
// 	validator Validator,
// 	repo OrderRepository,
// ) (Order, error) {

// 	valid := validator.CheckUser(
// 		req.UserID,
// 		token,
// 	)

// 	if !valid {
// 		return Order{}, errors.New(
// 			"user tidak valid",
// 		)
// 	}

// 	if req.Berat <= 0 {
// 		return Order{}, errors.New(
// 			"berat tidak valid",
// 		)
// 	}

// 	req.Resi = GenerateResi()
// 	req.ETA = CalculateETA()
// 	req.Status = "created"

// 	err := repo.Save(req)

// 	if err != nil {
// 		return Order{}, err
// 	}

// 	return req, nil
// }

// func GetOrder(id int) *Order {

// 	repo := MySQLRepository{}

// 	order, err := repo.GetByID(id)

// 	if err != nil {
// 		return nil
// 	}

// 	return order
// }

// // Tambahkan fungsi ini di bagian bawah service.go
// func GetAllOrders() []Order {
// 	repo := MySQLRepository{}
// 	orders, err := repo.GetAll()
// 	if err != nil {
// 		return []Order{}
// 	}
// 	return orders
// }

// func UpdateOrderStatus(
// 	id int,
// 	status string,
// ) bool {

// 	repo := MySQLRepository{}

// 	err := repo.UpdateStatus(
// 		id,
// 		status,
// 	)

// 	return err == nil
// }

// func GetETA(id int) string {

// 	order := GetOrder(id)

// 	if order == nil {
// 		return ""
// 	}

// 	return order.ETA
// }

package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Docker Compose
var UserServiceURL = "http://user-service:8081"

// Local testing
// var UserServiceURL = "http://localhost:8081"

type Validator interface {
	CheckUser(userID int, token string) bool
}

type RealValidator struct{}

func (v RealValidator) CheckUser(userID int, token string) bool {

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("PROFILE ERROR:", err)
		return false
	}

	fmt.Println("PROFILE STATUS:", resp.StatusCode)

	return resp.StatusCode == 200
}

func GenerateResi() string {
	return fmt.Sprintf(
		"LNG-%d",
		time.Now().UnixNano(),
	)
}

func CalculateETA() string {
	return "2 days"
}

func CreateOrder(
	req Order,
	token string,
	validator Validator,
	repo OrderRepository,
) (Order, error) {

	valid := validator.CheckUser(
		req.UserID,
		token,
	)

	if !valid {
		return Order{}, errors.New(
			"user tidak valid",
		)
	}

	if req.Berat <= 0 {
		return Order{}, errors.New(
			"berat tidak valid",
		)
	}

	req.Resi = GenerateResi()
	req.ETA = CalculateETA()
	req.Status = "created"

	// 🔥 DIUBAH DI SINI: Ditambahkan tanda '&' sebelum req
	// agar struct 'req' dilempar sebagai pointer, sehingga nilai OrderID hasil generate MySQL bisa disimpan masuk ke dalamnya.
	err := repo.Save(&req)

	if err != nil {
		return Order{}, err
	}

	return req, nil
}

func GetOrder(id int) *Order {

	repo := MySQLRepository{}

	order, err := repo.GetByID(id)

	if err != nil {
		return nil
	}

	return order
}

// Tambahkan fungsi ini di bagian bawah service.go
func GetAllOrders() []Order {
	repo := MySQLRepository{}
	orders, err := repo.GetAll()
	if err != nil {
		return []Order{}
	}
	return orders
}

func UpdateOrderStatus(
	id int,
	status string,
) bool {

	repo := MySQLRepository{}

	err := repo.UpdateStatus(
		id,
		status,
	)

	return err == nil
}

func GetETA(id int) string {

	order := GetOrder(id)

	if order == nil {
		return ""
	}

	return order.ETA
}
