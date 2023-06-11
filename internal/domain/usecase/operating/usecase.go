package operating

import (
	"context"
	"logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/room"
)

// Смысл юзкейса собрать все операции, которые носят только служебный характер.
// Никаких доменный сущностей он затрагивать не должен.
type Usecase struct {
	room *room.Usecase
	auth *auth.Usecase
}

func NewUsecase(room *room.Usecase, auth *auth.Usecase) *Usecase {
	return &Usecase{
		room: room,
		auth: auth,
	}
}

func (u *Usecase) UpdateSubscribes(ctx context.Context) error {
	if err := u.room.SubscribesRooms(ctx); err != nil {
		return err
	}
	//todo notifications channel
	return nil
}
