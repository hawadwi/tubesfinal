// package main

// type OrderRepository interface {
// 	Save(order Order) error
// }

// package main

// type OrderRepository interface {
// 	Save(order Order) error
// 	GetByID(id int) (*Order, error)
// 	UpdateStatus(id int, status string) error
// }

package main

type OrderRepository interface {
	Save(order *Order) error
	GetByID(id int) (*Order, error)
	GetAll() ([]Order, error) // <-- TAMBAHKAN BARIS INI
	UpdateStatus(id int, status string) error
}
