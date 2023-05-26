package app

import (
	"log"
	"logity/config"
	authUsecase "logity/internal/domain/usecase/auth"
	"logity/internal/infrustructure/repository/auth"
	"logity/internal/infrustructure/tokenManager"
	"logity/pkg/postgres"
)

func Run(cfg *config.Config) {
	dbClient, err := postgres.New(&cfg.Database)
	if err != nil {
		log.Fatalf("error init client db: %s", err)
	}

	authRepo := auth.NewUserRepository(dbClient)

	tokenMng := tokenManager.NewTokenManager(cfg)

	authUsecase.NewUserUsecase(authRepo, tokenMng)

}
