package dto

import validation "github.com/go-ozzo/ozzo-validation"

type SignInInputDto struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}

func (s SignInInputDto) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Login, validation.Required, validation.Length(4, 50)),
		validation.Field(&s.Password, validation.Required, validation.Length(4, 50)),
	)
}

type SignInOutputDto struct {
	AccessToken  JWT `json:"access_token" bson:"login"`
	RefreshToken JWT `json:"refresh_token" bson:"password"`
}
