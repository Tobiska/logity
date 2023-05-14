package user

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
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
	UserId   *uuid.UUID
	Username string
	Email    Email
	Phone    Phone
	Fio      string
}

func NewUser(username, email, phone, fio string) (*User, error) {
	pn, err := NewPhone(phone)
	if err != nil {
		return nil, err
	}

	m, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Email:    m,
		Phone:    pn,
		Fio:      fio,
	}, nil
}
