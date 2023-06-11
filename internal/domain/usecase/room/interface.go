package room

import (
	"context"
	"logity/internal/domain/entity/log"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto"
)

type (
	Repository interface {
		CreateRoom(ctx context.Context, userId string, room *room.Room) (*room.Room, error)
		GetRoomByCode(ctx context.Context, roomCode string) (*room.Room, error)
		UpdateRoom(ctx context.Context, dto dto.UpdateRoomDto) (*room.Room, error)
		DeleteRoom(ctx context.Context, roomCode string) (*room.Room, error)

		FindRoomByFilter(ctx context.Context, filter dto.FindFilter) ([]*room.Room, error)
		ShowAllCreatedRoom(ctx context.Context, userId string) ([]*room.Room, error)
		ShowAllAttachedRoom(ctx context.Context, userId string) ([]*room.Room, error)

		AttachUserToRoom(ctx context.Context, userId, roomCode string) (*room.Room, error)
		CheckInvite(ctx context.Context, userId, roomCode string) error
		InviteUserToRoom(ctx context.Context, userId, roomCode string) error
		DetachUserFromRoom(ctx context.Context, userId, roomCode string) error
	}

	Publisher interface {
		SubscribeUserOnRoom(ctx context.Context, u *user.User, r *room.Room) error
		UserRoomsUpdatedPublish(ctx context.Context, u *user.User, rs []*room.Room) error
		RoomUpdatedPublish(ctx context.Context, r *room.Room) error
		SendLogToRoomPublish(ctx context.Context, roomId string, log *log.Log) error
	}
)
