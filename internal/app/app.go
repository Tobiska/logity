package app

import (
	"fmt"
	"log"
	"logity/config"
	"logity/internal/delivery/rest"
	"logity/internal/domain/usecase"
	authUsecase "logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/room"
	"logity/internal/infrustructure/genHash/bcrypt"
	"logity/internal/infrustructure/repository/auth"
	room2 "logity/internal/infrustructure/repository/room"
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

	roomRepo := room2.NewRoomRepository(neo4jDriver, &cfg.Neo4j)

	userRepo := users.NewRepository(neo4jDriver, &cfg.Neo4j)

	tokenMng := tokenManager.NewTokenManager(cfg)

	authUc := authUsecase.NewUserUsecase(authRepo, userRepo, tokenMng)
	roomUc := room.NewRoomUsecase(roomRepo)

	env := usecase.NewEnv(roomUc, authUc)

	r := rest.NewRouter()
	rest.RegisterRouting(r, env)

	if err := http.ListenAndServe(cfg.ApiPort, r); err != nil {
		log.Fatalf("error http server %s", err)
	}
}
