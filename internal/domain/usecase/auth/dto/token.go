package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type UpdateTokenInputDto struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

func (d UpdateTokenInputDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.RefreshToken, validation.Required),
	)
}

type JWT struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
