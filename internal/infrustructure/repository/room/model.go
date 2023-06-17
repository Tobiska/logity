package room

import (
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
)

type Room struct {
	Id       string `json:"id"`
	Name     string `json:"username"`
	Tag      string `json:"tag"`
	Owner    User   `json:"owner"`
	Members  []User `json:"users"`
	Inviters []User `json:"invited"`
}

type User struct {
	Id       string  `json:"id"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Username string  `json:"username"`
}

func (r Room) toDomain() *room.Room {
	members := make([]*user.User, 0, len(r.Members))
	for _, u := range r.Members {
		members = append(members, u.toDomain())
	}

	inviters := make([]*user.User, 0, len(r.Inviters))
	for _, u := range r.Inviters {
		inviters = append(inviters, u.toDomain())
	}
	return &room.Room{
		Id:       r.Id,
		Name:     r.Name,
		Tag:      r.Tag,
		Owner:    r.Owner.toDomain(),
		Members:  members,
		Inviters: inviters,
	}

	//todo add log history
}

func (u User) toDomain() *user.User {
	return &user.User{
		Id:       u.Id,
		Username: u.Username,
		Phone:    (*user.Phone)(u.Phone),
		Email:    (*user.Email)(u.Email),
	}
}
