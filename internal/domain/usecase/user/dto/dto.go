package dto

import "logity/internal/domain/entity/user"

type (
	SignInInputDto struct {
		Login    string
		Password string
	}

	SignInOutputDto struct {
		AccessToken  string
		RefreshToken string
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
