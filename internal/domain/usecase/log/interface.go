package log

import (
	"context"
	"logity/internal/domain/entity/log"
)

type (
	Repository interface {
		CreateLogText(ctx context.Context, l *log.LogText, roomIds []string) error
		CreateLogPhoto(ctx context.Context, l *log.LogPhoto, roomIds []string) error
		CreateLogPicture(ctx context.Context, l *log.LogPicture, roomIds []string) error

		GetLogs(ctx context.Context) ([]*log.Log, error)
	}
	Publisher interface {
		PublishLogText(ctx context.Context, l *log.LogText, roomIds []string) error
		PublishLogPhoto(ctx context.Context, l *log.LogPhoto, roomIds []string) error
		PublishLogPicture(ctx context.Context, l *log.LogPicture, roomIds []string) error
	}
)
