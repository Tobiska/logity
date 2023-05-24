package auth

import (
	"context"
	"github.com/google/uuid"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/user/dto"
	"logity/internal/infrustructure/repository"
)

//Repository interface {
//CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error)
//CreateUser(ctx context.Context, u user.User, passwordHash string) (*user.User, error)
//SaveRefreshToken(ctx context.Context, u *user.User, refreshToken string) error
//CheckRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error)
//ResetPassword(ctx context.Context, dto dto.ResetPasswordDto) error
//}

type Repository struct {
	client repository.Client
}

func NewUserRepository(client repository.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error) {
	panic("imlement me!!")
}

func (r *Repository) CreateUser(ctx context.Context, u user.User, passwordHash string) (*user.User, error) {
	panic("imlement me!!")
}

func (r *Repository) SaveRefreshToken(ctx context.Context, u *user.User, refreshToken string) error {
	panic("imlement me!!")
}

func (r *Repository) CheckRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error) {
	panic("imlement me!!")
}

func (r *Repository) ResetPassword(ctx context.Context, dto dto.ResetPasswordDto) error {
	panic("imlement me!!")
}
