package app

import (
	"fmt"
	"log"
	"logity/config"
	"logity/internal/delivery/rest"
	authHanlder "logity/internal/delivery/rest/handlers/auth"
	authUsecase "logity/internal/domain/usecase/auth"
	"logity/internal/infrustructure/genHash/bcrypt"
	"logity/internal/infrustructure/repository/auth"
	"logity/internal/infrustructure/repository/users"
	"logity/internal/infrustructure/tokenManager"
	"logity/pkg/neo4j"
	"logity/pkg/postgres"
	"net/http"
)

func Run(cfg *config.Config) {
	dbClient, err := postgres.New(&cfg.Database)
	if err != nil {
		log.Fatalf("error init client db: %s", err)
	}

	neo4jDriver, err := neo4j.NewDriverNeo4j(&cfg.Neo4j)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error liquibase driver init: %s", err))
	}

	generator := bcrypt.NewGenerator(&cfg.Auth)

	authRepo := auth.NewUserRepository(dbClient, generator)

	userRepo := users.NewRepository(neo4jDriver, &cfg.Neo4j)

	tokenMng := tokenManager.NewTokenManager(cfg)

	authUc := authUsecase.NewUserUsecase(authRepo, userRepo, tokenMng)

	r := rest.NewRouter()
	authHanlder.NewHandler(authUc).Register(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error http server %s", err)
	}
}
