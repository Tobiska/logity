package room

import (
	"context"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto"
)

type (
	RoomRepository interface {
		CreateRoom(ctx context.Context, user *user.User, room *room.Room) error
		GetRoomByCode(ctx context.Context, roomCode string) (*room.Room, error)
		UpdateRoom(ctx context.Context, dto dto.UpdateRoomDto) (*room.Room, error)
		DeleteRoom(ctx context.Context, roomCode string) (*room.Room, error)

		FindRoomByFilter(ctx context.Context, filter dto.FindFilter) ([]*room.Room, error)
		ShowAllCreatedRoom(ctx context.Context, userId string) ([]*room.Room, error)
		ShowAllAttachedRoom(ctx context.Context, userId string) ([]*room.Room, error)

		AttachUserToRoom(ctx context.Context, userId, roomCode string) (*room.Room, error)
		DetachUserFromRoom(ctx context.Context, userId, roomCode string) (*room.Room, error)
	}
)
