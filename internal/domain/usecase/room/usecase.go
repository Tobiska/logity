package room

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto/input"
)

type Usecase struct {
	repo      Repository
	publisher Publisher
}

func NewRoomUsecase(repo Repository, publisher Publisher) *Usecase {
	return &Usecase{
		repo:      repo,
		publisher: publisher,
	}
}

func (us *Usecase) CreateNewRoom(ctx context.Context, dto input.CreateRoomDto) (*room.Room, error) {
	//todo реализовать транизакции: если комната создана, но owner на неё не подписан нужен rollback
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

	if err := us.publisher.SubscribeUserOnRoom(ctx, u, newRoom); err != nil {
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

	//todo оправить notifications

	return nil
}

func (us *Usecase) ShowRooms(ctx context.Context) ([]*room.Room, error) {
	u := user.ExtractFromCtx(ctx)
	if u == nil {
		return nil, fmt.Errorf("can't find the user")
	}

	rooms, err := us.repo.ShowAllAttachedRoom(ctx, u.Id)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (us *Usecase) SubscribesRooms(ctx context.Context) error {
	u := user.ExtractFromCtx(ctx)
	rooms, err := us.repo.ShowAllAttachedRoom(ctx, u.Id)
	if err != nil {
		return err
	}

	for _, r := range rooms {
		if err := us.publisher.SubscribeUserOnRoom(ctx, u, r); err != nil {
			fmt.Printf("error subscribing on room: %s\n", err) //todo убрать, когда появяться логи
		}
	}
	return nil
}

func (us *Usecase) JoinToRoom(ctx context.Context, dto input.JoinToRoomDto) (*room.Room, error) {
	//todo реализовать транизакции: если комната создана, но owner на неё не подписан нужен rollback
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

	if err := us.publisher.SubscribeUserOnRoom(ctx, u, r); err != nil {
		return nil, err
	}

	if err := us.publisher.RoomUpdatedPublish(ctx, r); err != nil {
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
