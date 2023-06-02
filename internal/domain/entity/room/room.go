package room

import (
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/user"
)

type Room struct {
	Id         string // служеный идентификатор
	Code       string // code для инвайта (уникальный)
	Name       string // имя комнаты для идентификации юзерами
	Tag        string // краткий тэг генерирующийся из Name
	Users      []*user.User
	LogHistory []*log.Log // последние 1000 логов
}
