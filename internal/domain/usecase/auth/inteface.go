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
	AuthRepository interface {
		CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error)
		CreateUser(ctx context.Context, d dto.CreateUserDto) (*user.User, error)
		SaveRefreshToken(ctx context.Context, u *user.User, refreshToken dto.JWT) error
		CheckRefreshToken(ctx context.Context, userId string, refreshToken dto.JWT) error
		ResetPassword(ctx context.Context, dto dto.ResetPasswordDto) error
		FindUser(ctx context.Context, userId string) (*user.User, error)
	}

	UserRepository interface {
		CreateUser(ctx context.Context, u *user.User) error
	}

	TokenManager interface {
		NewAccessToken(userId string) (dto.JWT, error)
		NewRefreshToken(userId string) (dto.JWT, error)
		NewRealTimeToken(userId string) (jwtToken dto.JWT, err error)
		ParseToken(accessToken string) (*dto.PayloadToken, error)
	}
)
