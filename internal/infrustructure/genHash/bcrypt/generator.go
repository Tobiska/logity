package bcrypt

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"logity/config"
	"logity/internal/domain/usecase/auth"
)

type Generator struct {
	cost int
}

func NewGenerator(cfg *config.Auth) *Generator {
	return &Generator{
		cost: cfg.PasswordHashCost,
	}
}

func (g Generator) Hash(_ context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), g.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (g Generator) Check(_ context.Context, raw, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return auth.ErrMismatch
	}
	return nil
}
