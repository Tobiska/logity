package output

type CreateRoomOutputDto struct {
	Room RoomOutputDto `json:"room"`
}

type RoomOutputDto struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	Tag      string           `json:"tag"`
	Members  []UsersOutputDto `json:"members"`
	Inviters []UsersOutputDto `json:"inviters"`
}

type UsersOutputDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
