package room

import (
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
)

type Room struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Tag   string `json:"tag"`
	Owner User   `json:"owner"`
	Users []User `json:"users"`
}

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}

func (r Room) toDomain() *room.Room {
	users := make([]*user.User, 0, len(r.Users))
	for _, u := range r.Users {
		users = append(users, u.toDomain())
	}
	return &room.Room{
		Id:    r.Id,
		Name:  r.Name,
		Tag:   r.Tag,
		Users: users,
	}

	//todo add log history
}

func (u User) toDomain() *user.User {
	return &user.User{
		Id:       u.Id,
		Username: u.Username,
		Phone:    user.Phone(u.Phone),
		Email:    user.Email(u.Email),
	}
}
