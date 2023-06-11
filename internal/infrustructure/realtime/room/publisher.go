package room

import (
	"context"
	"encoding/json"
	"fmt"
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
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

func (p *Publisher) SubscribeUserOnRoom(ctx context.Context, u *user.User, r *room.Room) error {
	channelName := realtime.GenerateRoomChannel(r.Id)
	if err := p.client.Subscribe(ctx, realtime.GenerateRoomChannel(r.Id), u.Id); err != nil {
		return fmt.Errorf("subscribe to channel: %s, user with id: %s with error: %w", channelName, u.Id, err)
	}
	return nil
}

func (p *Publisher) UserRoomsUpdatedPublish(ctx context.Context, u *user.User, rs []*room.Room) error {
	dto := NewRoomsDto(rs)
	msg, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("error marshal rooms message: %w", err)
	}
	channelName := realtime.GenerateRoomsUserChannel(u.Id)
	if _, err = p.client.Publish(ctx, channelName, msg); err != nil {
		return fmt.Errorf("publish to channel: %s, user with id: %s with error: %w", channelName, u.Id, err)
	}
	return nil
}

func (p *Publisher) RoomUpdatedPublish(ctx context.Context, r *room.Room) error {
	dto := NewRoomUpdatedDto(r)
	msg, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("error marshal room update message")
	}
	channelName := realtime.GenerateRoomChannel(r.Id)
	if _, err := p.client.Publish(ctx, channelName, msg); err != nil {
		return fmt.Errorf("publish to channel: %s, room with id: %s with error: %w", channelName, r.Id, err)
	}
	return nil
}

func (p *Publisher) SendLogToRoomPublish(ctx context.Context, roomId string, log *log.Log) error {
	return nil
}
