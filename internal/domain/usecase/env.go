package usecase

import (
	authUsecase "logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/room"
)

type Env struct {
	RoomUsecase *room.Usecase
	AuthUsecase *authUsecase.AuthUsecase
}

func NewEnv(roomUsecase *room.Usecase, authUsecase *authUsecase.AuthUsecase) *Env {
	return &Env{
		RoomUsecase: roomUsecase,
		AuthUsecase: authUsecase,
	}
}
