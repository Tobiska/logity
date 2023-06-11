package log

import (
	"logity/internal/domain/entity/user"
	"time"
)

type LogType string

const (
	TextType = "text"
)

type Log struct {
	Id        string
	Owner     *user.User
	Type      LogType
	CreatedAt time.Time
}

type LogText struct {
	Log
	Text string
}

func NewLogText(u *user.User, text string) *LogText {
	return &LogText{
		Log: Log{
			Owner:     u,
			Type:      TextType,
			CreatedAt: time.Now(),
		},
		Text: text,
	}
}

// todo возм. содержать мн-во характеристик(разрешение, размеры мета информация)
type LogPhoto struct {
	Log
	FilePath string //mb url
	Content  []byte
	Width    int
	Height   int
	Size     float64
}

type LogPicture struct {
	Log
	FilePath string //mb url
	Content  []byte
	Width    int
	Height   int
	Size     float64
}
