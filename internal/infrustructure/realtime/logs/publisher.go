package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"logity/internal/domain/entity/log"
	"logity/internal/infrustructure/realtime"
)

type Publisher struct {
	client realtime.Client
}

func NewPublisher(client realtime.Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) PublishLogText(ctx context.Context, l *log.LogText, roomIds []string) error {
	dto := NewLogDto(l)
	msg, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("error log message marshal: %w", err)
	}
	for _, id := range roomIds { //todo broadcast
		_, err := p.client.Publish(ctx, realtime.GenerateRoomChannel(id), msg)
		fmt.Printf("error publish %s\n", err)
	}
	return nil
}

func (p *Publisher) PublishLogPhoto(ctx context.Context, l *log.LogPhoto, roomIds []string) error {
	panic("imlement PublishLogPhoto")
}
func (p *Publisher) PublishLogPicture(ctx context.Context, l *log.LogPicture, roomIds []string) error {
	panic("imlement PublishLogPicture")
}
