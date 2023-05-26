package dto

import "time"

type JWT struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
