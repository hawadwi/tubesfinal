package main

import "testing"

func TestRegister(t *testing.T) {
	_, err := Register("Ula", "ula@mail.com", "123", "customer")
	if err != nil {
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	Register("Ula2", "ula2@mail.com", "123", "admin")
	_, err := Login("ula2@mail.com", "123")
	if err != nil {
		t.Fail()
	}
}

func TestUpdateProfile(t *testing.T) {

	u, err := Register(
		"A",
		"a@mail.com",
		"123",
		"customer",
	)

	if err != nil {
		t.Fatal(err)
	}

	ok := UpdateProfile(
		u.UserID,
		"Bandung",
		"Fast Delivery",
	)

	if !ok {
		t.Fatal("UpdateProfile returned false")
	}
}
