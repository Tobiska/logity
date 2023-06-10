package input

import validation "github.com/go-ozzo/ozzo-validation"

type InviteRoom struct {
	RoomId string `json:"room_id"`
	UserId string `json:"user_id"`
}

func (i *InviteRoom) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.RoomId, validation.Required, validation.Length(5, 100)),
		validation.Field(&i.UserId, validation.Required, validation.Length(5, 100)),
	)
}
