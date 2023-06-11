package log

import (
	"context"
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/log/dto/input"
)

type Usecase struct {
	repo      Repository
	publisher Publisher
}

func NewUsecase(repo Repository, publisher Publisher) *Usecase {
	return &Usecase{
		repo:      repo,
		publisher: publisher,
	}
}

func (ul *Usecase) PushTextLog(ctx context.Context, dto input.PushLogTextDto) error {
	u := user.ExtractFromCtx(ctx)
	l := log.NewLogText(u, dto.Text)

	if err := ul.repo.CreateLogText(ctx, l, dto.RoomIds); err != nil {
		return err
	}

	if err := ul.publisher.PublishLogText(ctx, l, dto.RoomIds); err != nil {
		return err
	}

	return nil
}

func (ul *Usecase) PushPhotoLog(ctx context.Context, dto input.PushLogPhotoDto) error {
	return nil
}

func (ul *Usecase) PushPictureLog(ctx context.Context, dto input.PushLogPictureDto) error {
	return nil
}
