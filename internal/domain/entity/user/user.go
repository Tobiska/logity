package user

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var phoneRegex = regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)

type Email string

func NewEmail(s string) (Email, error) {
	if !emailRegex.Match([]byte(s)) {
		return "", fmt.Errorf("email validation error")
	}
	return Email(s), nil
}

type Phone string

func NewPhone(s string) (Phone, error) {
	if !phoneRegex.Match([]byte(s)) {
		return "", fmt.Errorf("phone validation error")
	}
	return Phone(s), nil
}

type User struct {
	Id           string
	Email        Email
	Phone        Phone
	Username     string
	PasswordHash string
}

func NewUser(email, phone, fio string) (*User, error) {
	pn, err := NewPhone(phone)
	if err != nil {
		return nil, err
	}

	m, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    m,
		Phone:    pn,
		Username: fio,
	}, nil
}
