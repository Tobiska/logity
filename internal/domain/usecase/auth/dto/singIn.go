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
	AccessToken  JWT    `json:"access_token" bson:"access_token"`
	RefreshToken JWT    `json:"refresh_token" bson:"refresh_token"`
	RTCToken     JWT    `json:"rtc_token" bson:"rtc_token"`
	RTCHost      string `json:"rtc_host" bson:"rtc_host"`
}
