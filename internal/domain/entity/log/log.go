package log

import (
	"logity/internal/domain/entity/user"
	"time"
)

type Log struct {
	Id        string
	U         *user.User
	Cnt       IContent
	CreatedAt time.Time
}
