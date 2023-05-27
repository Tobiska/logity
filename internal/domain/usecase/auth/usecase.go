package auth

import (
	"context"
	"fmt"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/auth/dto"
)

type AuthUsecase struct {
	repo         Repository
	tokenManager TokenManager
}

func NewUserUsecase(repo Repository, manager TokenManager) *AuthUsecase {
	return &AuthUsecase{
		repo:         repo,
		tokenManager: manager,
	}
}

func (us AuthUsecase) RevokeRefreshToken(_ context.Context, _ string) error {
	panic("implement RevokeRefreshToken me!!!")
}

func (us AuthUsecase) FindUserByAccessToken(ctx context.Context, accessToken string) (*user.User, error) {
	payload, err := us.tokenManager.ParseToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("parse token error: %w", err)
	}

	u, err := us.repo.FindUser(ctx, payload.UserId)
	if err != nil {
		return nil, fmt.Errorf("find user error: %w", err)
	}

	return u, err
}

func (us AuthUsecase) Me(ctx context.Context) (*dto.MeOutputDto, error) {
	u := user.ExtractFromCtx(ctx)

	if u == nil {
		return nil, fmt.Errorf("error user doesn't contained context")
	}

	return &dto.MeOutputDto{
		UserId: u.Id,
		Email:  string(u.Email),
		Phone:  string(u.Phone),
		Fio:    u.Fio,
	}, nil
}

func (us AuthUsecase) SignIn(ctx context.Context, d dto.SignInInputDto) (*dto.SignInOutputDto, error) {
	u, err := us.repo.CheckCredentials(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error check credentials: %w", err)
	}

	if u.Id == "" {
		return nil, fmt.Errorf("error auth id is null")
	}
	userId := u.Id

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

func (us AuthUsecase) UpdateAccessToken(ctx context.Context, in dto.UpdateTokenInputDto) (jwtToken dto.JWT, err error) {
	payload, err := us.tokenManager.ParseToken(in.RefreshToken)
	if err != nil {
		return jwtToken, fmt.Errorf("error parse refresh token: %w", err)
	}

	if err := us.repo.CheckRefreshToken(ctx, payload.UserId, payload.Token); err != nil {
		return jwtToken, fmt.Errorf("error check refresh token: %w", err)
	}

	accessToken, err := us.tokenManager.NewJWT(payload.UserId)
	if err != nil {
		return jwtToken, fmt.Errorf("error generate access token %w", err)
	}

	return accessToken, nil
}

func (us AuthUsecase) SignUp(ctx context.Context, d dto.SignUpInputDto) (*dto.SignUpOutputDto, error) {
	u, err := us.repo.CreateUser(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error create auth: %w", err)
	}
	return &dto.SignUpOutputDto{
		UserId: u.Id,
	}, nil
}
