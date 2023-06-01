package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/auth/dto"
	"logity/internal/infrustructure/repository"
	"logity/pkg/postgres"
	"time"
)

const (
	UsersTable  = `users`
	TokensTable = `tokens`
)

type Repository struct {
	client        repository.Client
	hashGenerator repository.HashGenerator
}

func NewUserRepository(client repository.Client, hashGenerator repository.HashGenerator) *Repository {
	return &Repository{
		client:        client,
		hashGenerator: hashGenerator,
	}
}

func (r *Repository) CheckCredentials(ctx context.Context, dto dto.SignInInputDto) (*user.User, error) {
	//todo with pgcrypt check bcrypt hash in psql
	query, args, err := goqu.Dialect(postgres.Dialect).From(UsersTable).Prepared(true).Select(
		goqu.C(`id`), goqu.C(`email`), goqu.C(`phone`), goqu.C(`fio`), goqu.C(`password_hash`)).Where(
		goqu.ExOr{"phone": dto.Login, "email": dto.Login},
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("check credentials constuct query error: %w", err)
	}

	u := &user.User{}
	err = r.client.QueryRow(ctx, query, args...).Scan(&u.Id, &u.Email, &u.Phone, &u.Username, &u.PasswordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("auth with login: %s doesn't exist", dto.Login)
	}
	if err != nil {
		return nil, fmt.Errorf("exec query error check credentials: %w", err)
	}

	if err := r.hashGenerator.Check(ctx, dto.Password, u.PasswordHash); err != nil {
		return nil, fmt.Errorf("error check password %w", err)
	}

	return u, nil
}

func (r *Repository) FindUser(ctx context.Context, userId string) (*user.User, error) {
	query, args, err := goqu.Dialect(postgres.Dialect).From(goqu.T(UsersTable)).Prepared(true).
		Select(goqu.C("id"), goqu.C("email"), goqu.C("phone"), goqu.C("fio")).
		Where(goqu.Ex{"id": goqu.V(userId)}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("error contruct find user query: %w", err)
	}

	u := &user.User{}
	if err := r.client.QueryRow(ctx, query, args...).Scan(&u.Id, &u.Email, &u.Phone, &u.Username); err != nil {
		return nil, fmt.Errorf("error exec find user query: %w", err)
	}

	return u, nil
}

func (r *Repository) CreateUser(ctx context.Context, d dto.SignUpInputDto) (*user.User, error) {
	passwordHash, err := r.hashGenerator.Hash(ctx, d.Password)
	if err != nil {
		return nil, fmt.Errorf("password hash error: %w", err)
	}
	query, args, err := goqu.Dialect(postgres.Dialect).From(goqu.T(UsersTable)).Prepared(true).
		Insert().Rows(goqu.Record{
		"email":         goqu.V(d.Email),
		"phone":         goqu.V(d.Phone),
		"fio":           goqu.V(d.Fio),
		"password_hash": goqu.V(passwordHash),
	}).Returning(goqu.C("id")).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("create auth construct query error: %w", err)
	}

	var id string
	if err := r.client.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return nil, fmt.Errorf("exec error create auth error: %w", err)
	}

	u, err := user.NewUser(d.Email, d.Phone, d.Fio)
	if err != nil {
		return nil, fmt.Errorf("error create domain object user: %w", err)
	}

	u.Id = id

	return u, nil
}

func (r *Repository) SaveRefreshToken(ctx context.Context, u *user.User, token dto.JWT) error {
	query, args, err := goqu.Dialect(postgres.Dialect).From(TokensTable).Insert().Rows(goqu.Record{
		"user_id":       goqu.V(u.Id),
		"refresh_token": goqu.V(token.Token),
		"expired_at":    goqu.V(token.ExpiredAt),
	}).Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("save refresh token construct query error: %w", err)
	}

	if _, err := r.client.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("error save refresh token exec: %w", err)
	}

	return nil
}

func (r *Repository) CheckRefreshToken(ctx context.Context, userId string, refreshToken dto.JWT) error {
	query, args, err := goqu.Dialect(postgres.Dialect).From(TokensTable).Select(goqu.C("expired_at")).Prepared(true).
		Where(goqu.Ex{
			"user_id":       goqu.V(userId),
			"refresh_token": goqu.V(refreshToken.Token),
		}).ToSQL()
	if err != nil {
		return fmt.Errorf("save refresh token construct query error: %w", err)
	}

	var expiredAt time.Time
	if err := r.client.QueryRow(ctx, query, args...).Scan(&expiredAt); err != nil {
		return fmt.Errorf("error save refresh token exec: %w", err)
	}

	if time.Now().After(expiredAt) {
		return auth.ErrRefreshTokenExpired
	}

	return nil
}

func (r *Repository) ResetPassword(_ context.Context, _ dto.ResetPasswordDto) error {
	//todo implement
	panic("imlement me!!")
}
