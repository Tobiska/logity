package dto

import "logity/internal/domain/entity/user"

type (
	SignInInputDto struct {
		Login    string
		Password string
	}

	SignInOutputDto struct {
		AccessToken  JWT `json:"access_token"`
		RefreshToken JWT `json:"refresh_token"`
	}

	SignUpInputDto struct {
		user.User
		Password string
	}

	ResetPasswordDto struct {
		UserId   string
		Password string
	}

	PayloadToken struct {
		UserId string
	}
)
