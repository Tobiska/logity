package dto

import (
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
)

type RoomOutputDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Owner    UserOutputDto
	Members  []UserOutputDto `json:"members"`
	Inviters []UserOutputDto `json:"inviters"`
}

type UserOutputDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewRoomOutputDto(r *room.Room) RoomOutputDto {
	members := make([]UserOutputDto, 0)
	for _, u := range r.Members {
		members = append(members, NewUserOutputDto(u))
	}

	inviters := make([]UserOutputDto, 0)
	for _, u := range r.Inviters {
		inviters = append(inviters, NewUserOutputDto(u))
	}
	return RoomOutputDto{
		Id:       r.Id,
		Name:     r.Name,
		Tag:      r.Tag,
		Owner:    NewUserOutputDto(r.Owner),
		Members:  members,
		Inviters: inviters,
	}
}

func NewUserOutputDto(u *user.User) UserOutputDto {
	return UserOutputDto{
		Id:    u.Id,
		Email: string(u.Email),
		Phone: string(u.Phone),
	}
}
