package room

import validation "github.com/go-ozzo/ozzo-validation"

type CreateRoom struct {
	Name string `json:"name"`
}

func (c *CreateRoom) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(5, 100)),
	)
}
