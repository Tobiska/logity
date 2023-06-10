package room

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto/input"
)

type Usecase struct {
	repo RoomRepository
}

func NewRoomUsecase(repo RoomRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (us *Usecase) CreateNewRoom(ctx context.Context, dto input.CreateRoomDto) (*room.Room, error) {
	u := user.ExtractFromCtx(ctx)
	if u == nil {
		return nil, fmt.Errorf("can't find the user")
	}

	createdRoom, err := us.repo.CreateRoom(ctx, u.Id, room.NewFromRoomName(dto.Name))
	if err != nil {
		return nil, err
	}

	newRoom, err := us.repo.GetRoomByCode(ctx, createdRoom.Id)
	if err != nil {
		return nil, err
	}

	return newRoom, nil
}

func (us *Usecase) InviteToRoom(ctx context.Context, dto input.InviteToRoomDto) error {
	u := user.ExtractFromCtx(ctx)
	if u == nil {
		return fmt.Errorf("can't find the user")
	}

	if err := us.repo.InviteUserToRoom(ctx, dto.UserId, dto.RoomId); err != nil {
		return err
	}

	return nil
}

func (us *Usecase) JoinToRoom(ctx context.Context, dto input.JoinToRoomDto) (*room.Room, error) {
	u := user.ExtractFromCtx(ctx)
	if u == nil {
		return nil, fmt.Errorf("can't find the user")
	}

	if err := us.repo.CheckInvite(ctx, u.Id, dto.RoomId); err != nil {
		return nil, err
	}

	r, err := us.repo.AttachUserToRoom(ctx, u.Id, dto.RoomId)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (us *Usecase) GetRoom(ctx context.Context, roomId string) (*room.Room, error) {
	r, err := us.repo.GetRoomByCode(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return r, nil
}
