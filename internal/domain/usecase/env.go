package usecase

import (
	authUsecase "logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/log"
	"logity/internal/domain/usecase/operating"
	"logity/internal/domain/usecase/room"
)

type Env struct {
	RoomUsecase      *room.Usecase
	AuthUsecase      *authUsecase.Usecase
	OperatingUsecase *operating.Usecase
	LogUsecase       *log.Usecase
}

func NewEnv(roomUsecase *room.Usecase, authUsecase *authUsecase.Usecase, operatingUsecase *operating.Usecase, logUsecase *log.Usecase) *Env {
	return &Env{
		RoomUsecase:      roomUsecase,
		AuthUsecase:      authUsecase,
		OperatingUsecase: operatingUsecase,
		LogUsecase:       logUsecase,
	}
}
