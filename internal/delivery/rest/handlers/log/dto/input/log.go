package input

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LogTextInput struct {
	Text    string   `json:"text"`
	RoomIds []string `json:"room_ids"`
}

func (i *LogTextInput) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Text, validation.Required, validation.Length(1, 255)),
		validation.Field(&i.RoomIds, validation.Required, validation.Length(1, 255), validation.Each(is.UUID)),
	)
}
