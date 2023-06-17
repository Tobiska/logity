package logs

import (
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/user"
	"time"
)

type LogDto struct {
	Id        string  `json:"id"`
	Type      string  `json:"type"`
	User      UserDto `json:"user"`
	Text      string  `json:"text"`
	CreatedAt string  `json:"created_at"`
}

func NewLogDto(l *log.LogText) LogDto {
	return LogDto{
		Id:        l.Id,
		Type:      string(l.Type),
		User:      NewUserDto(l.Owner),
		Text:      l.Text,
		CreatedAt: l.CreatedAt.Format(time.RFC3339),
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
