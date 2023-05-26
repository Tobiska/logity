package auth

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/auth/dto"
)

var (
	ErrMismatch            = fmt.Errorf("error mismatch hash and password")
	ErrRefreshTokenExpired = fmt.Errorf("error refresh token expired")
)

type (
	Repository interface {
		CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error)
		CreateUser(ctx context.Context, u user.User) (*user.User, error)
		SaveRefreshToken(ctx context.Context, u *user.User, refreshToken dto.JWT) error
		CheckRefreshToken(ctx context.Context, userId string, refreshToken dto.JWT) error
		ResetPassword(ctx context.Context, dto dto.ResetPasswordDto) error
	}

	TokenManager interface {
		NewJWT(userId string) (dto.JWT, error)
		NewRefreshToken(userId string) (dto.JWT, error)
		ParseToken(accessToken string) (*dto.PayloadToken, error)
	}
)
