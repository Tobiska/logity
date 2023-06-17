package dto

import "logity/internal/domain/entity/user"

type MeOutputDto struct {
	UserId string  `json:"user_id"`
	Email  *string `json:"email"`
	Phone  *string `json:"phone"`
	Fio    string  `json:"fio"`
}

func MeDtoFromDomain(u *user.User) *MeOutputDto {
	return &MeOutputDto{
		UserId: u.Id,
		Email:  (*string)(u.Email),
		Phone:  (*string)(u.Phone),
		Fio:    u.Username,
	}
}
