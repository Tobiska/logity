package output

import (
	"logity/internal/domain/entity/room"
)

type RoomShowOutputDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func NewRoomShowOutputDto(r *room.Room) RoomShowOutputDto {
	return RoomShowOutputDto{
		Id:   r.Id,
		Name: r.Name,
		Tag:  r.Tag,
	}
}

type ShowRoomsOutputDto struct {
	Rooms []RoomShowOutputDto `json:"rooms"`
}

func NewShowRoomsOutputDto(rs []*room.Room) ShowRoomsOutputDto {
	roomsDto := make([]RoomShowOutputDto, 0, len(rs))
	for _, r := range rs {
		roomsDto = append(roomsDto, NewRoomShowOutputDto(r))
	}
	return ShowRoomsOutputDto{
		Rooms: roomsDto,
	}
}
