package bcrypt

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"logity/config"
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
