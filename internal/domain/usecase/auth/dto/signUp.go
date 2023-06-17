package dto

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

var (
	ErrConfirmPassword = fmt.Errorf("configrm password error")
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var phoneRegex = regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)

// todo реализовать регистрацию через номер телефона
type SignUpByEmailInputDto struct {
	Email           string `json:"email" bson:"email"`
	Fio             string `json:"fio" bson:"fio"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
}

func (d SignUpByEmailInputDto) Validate() error {
	if err := validation.ValidateStruct(&d,
		validation.Field(&d.Email, validation.Required, validation.Match(emailRegex), validation.Length(10, 50)),
		validation.Field(&d.Fio, validation.Length(0, 100)),
		validation.Field(&d.Password, validation.Required, validation.Length(5, 200)),
		validation.Field(&d.ConfirmPassword, validation.Required, validation.Length(5, 200)),
	); err != nil {
		return err
	}

	if d.Password != d.ConfirmPassword {
		return ErrConfirmPassword
	}

	return nil
}

type SignUpOutputDto struct {
	UserId string `json:"user_id" bson:"user_id"`
}
