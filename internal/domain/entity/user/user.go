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
	Email        *Email
	Phone        *Phone
	Username     string
	PasswordHash string
}

func NewUser(email, phone *string, fio string) (*User, error) {
	u := &User{
		Username: fio,
	}

	if phone != nil {
		pn, err := NewPhone(*phone)
		if err != nil {
			return nil, err
		}
		u.Phone = &pn
	}

	if email != nil {
		m, err := NewEmail(*email)
		if err != nil {
			return nil, err
		}
		u.Email = &m
	}

	return u, nil
}
