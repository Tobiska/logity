package room

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto/input"
	"logity/internal/domain/usecase/room/dto/output"
)

type Usecase struct {
	repo RoomRepository
}

func NewRoomUsecase(repo RoomRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (us *Usecase) CreateNewRoom(ctx context.Context, dto input.CreateRoomDto) (*output.CreateRoomOutputDto, error) {
	u := user.ExtractFromCtx(ctx)
	if u == nil {
		return nil, fmt.Errorf("can't find the user")
	}

	r := room.NewFromRoomName(dto.Name)

	if err := us.repo.CreateRoom(ctx, u, r); err != nil {
		return nil, err
	}

	r, err := us.repo.AttachUserToRoom(ctx, u.Id, r.Code)
	if err != nil {
		return nil, err
	}

	return &output.CreateRoomOutputDto{
		Code:       r.Code,
		Name:       r.Name,
		Tag:        r.Tag,
		CountUsers: len(r.Users),
	}, nil
}
