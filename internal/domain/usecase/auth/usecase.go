package auth

import (
	"context"
	"fmt"
	"logity/config"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/auth/dto"
)

type Usecase struct {
	authRepo     AuthRepository
	tokenManager TokenManager
	userRepo     UserRepository
	cfg          *config.Config
}

func NewUserUsecase(repo AuthRepository, userRepo UserRepository, manager TokenManager, cfg *config.Config) *Usecase {
	return &Usecase{
		authRepo:     repo,
		tokenManager: manager,
		userRepo:     userRepo,
		cfg:          cfg,
	}
}

func (us Usecase) RevokeRefreshToken(_ context.Context, _ string) error {
	panic("implement RevokeRefreshToken me!!!")
}

func (us Usecase) FindUserByAccessToken(ctx context.Context, accessToken string) (*user.User, error) {
	payload, err := us.tokenManager.ParseToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("parse token error: %w", err)
	}

	u, err := us.authRepo.FindUser(ctx, payload.UserId)
	if err != nil {
		return nil, fmt.Errorf("find user error: %w", err)
	}

	return u, err
}

func (us Usecase) Me(ctx context.Context) (*dto.MeOutputDto, error) {
	u := user.ExtractFromCtx(ctx)

	if u == nil {
		return nil, fmt.Errorf("error user doesn't contained context")
	}

	return &dto.MeOutputDto{
		UserId: u.Id,
		Email:  string(u.Email),
		Phone:  string(u.Phone),
		Fio:    u.Username,
	}, nil
}

func (us Usecase) SignIn(ctx context.Context, d dto.SignInInputDto) (*dto.SignInOutputDto, error) {
	u, err := us.authRepo.CheckCredentials(ctx, d)
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

	if err := us.authRepo.SaveRefreshToken(ctx, u, refreshToken); err != nil {
		return nil, fmt.Errorf("error save refresh token: %w", err)
	}

	//todo после сохранения токена его сразу же захотят отозвать, тогда accessToken будет не действительный ровно до истечения.

	accessToken, err := us.tokenManager.NewAccessToken(userId)
	if err != nil {
		return nil, fmt.Errorf("error generate access token: %w", err)
	}

	rtcToken, err := us.tokenManager.NewRealTimeToken(u.Id)
	if err != nil {
		return nil, fmt.Errorf("error generate rtc token: %w", err)
	}

	return &dto.SignInOutputDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		RTCToken:     rtcToken,
		RTCHost:      us.cfg.ClientHost,
	}, nil
}

func (us Usecase) UpdateRealTimeServerToken(ctx context.Context) (dto.JWT, error) {
	u := user.ExtractFromCtx(ctx)
	token, err := us.tokenManager.NewRealTimeToken(u.Id)
	if err != nil {
		return dto.JWT{}, fmt.Errorf("error generate rtc token: %w", err)
	}

	return token, nil
}

func (us Usecase) UpdateAccessToken(ctx context.Context, in dto.UpdateTokenInputDto) (jwtToken dto.JWT, err error) {
	payload, err := us.tokenManager.ParseToken(in.RefreshToken)
	if err != nil {
		return jwtToken, fmt.Errorf("error parse refresh token: %w", err)
	}

	if err := us.authRepo.CheckRefreshToken(ctx, payload.UserId, payload.Token); err != nil {
		return jwtToken, fmt.Errorf("error check refresh token: %w", err)
	}

	accessToken, err := us.tokenManager.NewAccessToken(payload.UserId)
	if err != nil {
		return jwtToken, fmt.Errorf("error generate access token %w", err)
	}

	return accessToken, nil
}

func (us Usecase) SignUp(ctx context.Context, d dto.SignUpInputDto) (*dto.SignUpOutputDto, error) {
	u, err := us.authRepo.CreateUser(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error create auth: %w", err)
	}

	if err := us.userRepo.CreateUser(ctx, u); err != nil {
		return nil, fmt.Errorf("error create user: %w", err)
	}
	return &dto.SignUpOutputDto{
		UserId: u.Id,
	}, nil
}
