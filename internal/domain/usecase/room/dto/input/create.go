package input

type CreateRoomDto struct {
	Name string `json:"name"`
}

type InviteToRoomDto struct {
	UserId string `json:"user_id"`
	RoomId string `json:"room_id"`
}

type JoinToRoomDto struct {
	RoomId string `json:"room_id"`
}
