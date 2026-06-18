package main

import "errors"

var users []User
var nextID = 1

func init() {

	seed := User{
		UserID:   nextID,
		Name:     "seed",
		Email:    "seed@mail.com",
		Password: "hashed",
		Role:     "customer",
	}

	users = append(users, seed)
	nextID++
}

func Register(
	name string,
	email string,
	password string,
	role string,
) (User, error) {

	user := User{
		UserID:   nextID,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}

	nextID++

	users = append(users, user)

	return user, nil
}

func Login(email string, password string) (User, error) {

	for _, u := range users {

		if u.Email == email &&
			u.Password == password {

			return u, nil
		}
	}

	return User{}, errors.New("login gagal")
}

func GetProfile(id int) *User {

	for i := range users {

		if users[i].UserID == id {
			return &users[i]
		}
	}

	return nil
}

func UpdateProfile(
	userID int,
	alamat string,
	preferensi string,
) bool {

	for i := range users {

		if users[i].UserID == userID {

			users[i].Alamat = alamat
			users[i].Preferensi = preferensi

			return true
		}
	}

	return false
}
