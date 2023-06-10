package output

import (
	"logity/internal/delivery/rest/handlers/room/dto"
	"logity/internal/domain/entity/room"
)

type JoinOutputDto struct {
	Room dto.RoomOutputDto `json:"room"`
}

func NewJoinOutputDto(r *room.Room) CreateRoomOutputDto {
	return CreateRoomOutputDto{
		Room: dto.NewRoomOutputDto(r),
	}
}
