package user

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/user/dto"
)

type User struct {
	repo          Repository
	tokenManager  TokenManager
	hashGenerator HashGenerator
}

func NewUserUsecase(repo Repository, manager TokenManager, hashGenerator HashGenerator) *User {
	return &User{
		repo:          repo,
		tokenManager:  manager,
		hashGenerator: hashGenerator,
	}
}

func (us User) SignIn(ctx context.Context, d dto.SignInInputDto) (*dto.SignInOutputDto, error) {
	u, err := us.repo.CheckCredentials(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error check credentials: %w", err)
	}

	if u.UserId == nil {
		return nil, fmt.Errorf("error user id is null")
	}
	userId := u.UserId.String()

	refreshToken, err := us.tokenManager.NewRefreshToken(userId)
	if err != nil {
		return nil, fmt.Errorf("error generate refresh token: %w", err)
	}

	if err := us.repo.SaveRefreshToken(ctx, u, refreshToken); err != nil {
		return nil, fmt.Errorf("error save refresh token: %w", err)
	}

	//todo после сохранения токена его сразу же закхотят отозвать, тогда accessToken будет не действительный ровно до истечения.

	accessToken, err := us.tokenManager.NewJWT(userId)
	if err != nil {
		return nil, fmt.Errorf("error generate access token: %w", err)
	}

	return &dto.SignInOutputDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (us User) UpdateAccessToken(ctx context.Context, refreshToken string) (string, error) {
	userId, err := us.repo.CheckRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("error check refresh token: %w", err)
	}

	accessToken, err := us.tokenManager.NewJWT(userId.String())
	if err != nil {
		return "", fmt.Errorf("error generate access token %w", err)
	}

	return accessToken, nil
}

func (us User) SignUp(ctx context.Context, d dto.SignUpInputDto) (*user.User, error) {
	h, err := us.hashGenerator.Hash(d.Password)
	if err != nil {
		return nil, fmt.Errorf("error hash from password %w", err)
	}
	u, err := us.repo.CreateUser(ctx, d.User, h)
	if err != nil {
		return nil, fmt.Errorf("error create user: %w", err)
	}
	return u, nil
}
