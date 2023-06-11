package realtime

import (
	"context"
	"github.com/centrifugal/gocent/v3"
)

type Client interface {
	Pipe() *gocent.Pipe
	Publish(ctx context.Context, channel string, data []byte, opts ...gocent.PublishOption) (gocent.PublishResult, error)
	Broadcast(ctx context.Context, channels []string, data []byte, opts ...gocent.PublishOption) (gocent.BroadcastResult, error)
	Subscribe(ctx context.Context, channel, user string, opts ...gocent.SubscribeOption) error
	Unsubscribe(ctx context.Context, channel, user string, opts ...gocent.UnsubscribeOption) error
	Disconnect(ctx context.Context, user string, opts ...gocent.DisconnectOption) error
	Presence(ctx context.Context, channel string) (gocent.PresenceResult, error)
	PresenceStats(ctx context.Context, channel string) (gocent.PresenceStatsResult, error)
	History(ctx context.Context, channel string, opts ...gocent.HistoryOption) (gocent.HistoryResult, error)
	HistoryRemove(ctx context.Context, channel string) error
	Channels(ctx context.Context, opts ...gocent.ChannelsOption) (gocent.ChannelsResult, error)
	Info(ctx context.Context) (gocent.InfoResult, error)
	SendPipe(ctx context.Context, pipe *gocent.Pipe) ([]gocent.Reply, error)
}
