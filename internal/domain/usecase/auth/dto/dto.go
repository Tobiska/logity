package dto

type (
	ResetPasswordDto struct {
		UserId   string
		Password string
	}

	PayloadToken struct {
		UserId string
		Token  JWT
	}
)
