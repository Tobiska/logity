package auth

import validation "github.com/go-ozzo/ozzo-validation"

type SignIn struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}

func (s SignIn) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Login, validation.Required, validation.Length(4, 50)),
		validation.Field(&s.Password, validation.Required, validation.Length(4, 50)),
	)
}
