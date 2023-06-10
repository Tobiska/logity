package room

import (
	"fmt"
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/user"
	"strings"
)

var (
	ErrTooShortRoomName = fmt.Errorf("too short room name")
)

type Room struct {
	Id         string // служеный идентификатор
	Name       string // имя комнаты для идентификации юзерами
	Tag        string // краткий тэг генерирующийся из Name
	Owner      *user.User
	Members    []*user.User
	Inviters   []*user.User
	LogHistory []*log.Log // последние 1000 логов
}

func NewFromRoomName(name string) *Room {
	return &Room{
		Name: name,
		Tag:  generateTag(name),
	}
}

func generateTag(str string) string { //todo придумать более изящный алгоритм
	return strings.ToUpper(string([]rune(str)[:4]))
}
