package dto

type (
	CreateUserDto struct {
		Email    *string
		Phone    *string
		Fio      string
		Password string
	}
)
