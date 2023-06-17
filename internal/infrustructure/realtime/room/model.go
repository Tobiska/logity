package room

import (
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
)

type RoomsUpdatedDto struct {
	Rooms []RoomShowDto `json:"rooms"`
}

func NewRoomsDto(rs []*room.Room) RoomsUpdatedDto {
	rooms := make([]RoomShowDto, 0, len(rs))
	for _, r := range rs {
		rooms = append(rooms, NewFromRoomDto(r))
	}
	return RoomsUpdatedDto{
		Rooms: rooms,
	}
}

type UserDto struct {
	Id    string  `json:"id"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func NewUserDto(u *user.User) UserDto {
	return UserDto{
		Id:    u.Id,
		Email: (*string)(u.Email),
		Phone: (*string)(u.Phone),
	}
}

type RoomShowDto struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Tag          string  `json:"tag"`
	Owner        UserDto `json:"owner"`
	MembersSize  int     `json:"members"`
	InvitersSize int     `json:"inviters"`
}

func NewFromRoomDto(r *room.Room) RoomShowDto {
	return RoomShowDto{
		Id:           r.Id,
		Name:         r.Name,
		Tag:          r.Tag,
		Owner:        NewUserDto(r.Owner),
		MembersSize:  len(r.Members),
		InvitersSize: len(r.Inviters),
	}
}

type RoomUpdatedDto struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Tag          string    `json:"tag"`
	Owner        UserDto   `json:"owner"`
	MembersSize  []UserDto `json:"members"`
	InvitersSize []UserDto `json:"inviters"`
	//todo logs
}

func NewRoomUpdatedDto(r *room.Room) RoomUpdatedDto {
	members := make([]UserDto, 0, len(r.Members))
	for _, m := range r.Members {
		members = append(members, NewUserDto(m))
	}
	inviters := make([]UserDto, 0, len(r.Inviters))
	for _, i := range r.Inviters {
		inviters = append(inviters, NewUserDto(i))
	}

	return RoomUpdatedDto{
		Id:           r.Id,
		Name:         r.Name,
		Tag:          r.Tag,
		Owner:        NewUserDto(r.Owner),
		MembersSize:  members,
		InvitersSize: inviters,
	}
}
