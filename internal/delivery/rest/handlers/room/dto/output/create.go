package output

import (
	"logity/internal/delivery/rest/handlers/room/dto"
	"logity/internal/domain/entity/room"
)

type CreateRoomOutputDto struct {
	Room dto.RoomOutputDto `json:"room"`
}

func NewCreateRoomOutputDto(r *room.Room) CreateRoomOutputDto {
	return CreateRoomOutputDto{
		Room: dto.NewRoomOutputDto(r),
	}
}
