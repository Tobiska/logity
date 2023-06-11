package cetrifugo

import (
	"context"
	"fmt"
	"github.com/centrifugal/gocent/v3"
	"logity/config"
)

func NewCentrifugo(cfg *config.Centrifugo) (*gocent.Client, error) {
	client := gocent.New(gocent.Config{
		Addr: cfg.ApiHost,
		Key:  cfg.ApiKey,
	})
	info, err := client.Info(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println(info)

	return client, nil
}
