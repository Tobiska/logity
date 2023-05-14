package user

import (
	"context"
	"github.com/google/uuid"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/user/dto"
)

type (
	Repository interface {
		CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error)
		CreateUser(ctx context.Context, u user.User, passwordHash string) (*user.User, error)
		SaveRefreshToken(ctx context.Context, u *user.User, refreshToken string) error
		CheckRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error)
		ResetPassword(ctx context.Context, dto dto.ResetPasswordDto) error
	}

	TokenManager interface {
		NewJWT(userId string) (string, error)
		NewRefreshToken(userId string) (string, error)
		ParseToken(accessToken string) (*dto.PayloadToken, error)
	}
	HashGenerator interface {
		Hash(ctx context.Context, raw string) (string, error)
	}
)
