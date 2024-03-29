package app

import (
	"context"
	"fmt"
	"log"
	"logity/config"
	"logity/internal/delivery/rest"
	"logity/internal/domain/usecase"
	authUsecase "logity/internal/domain/usecase/auth"
	log2 "logity/internal/domain/usecase/log"
	"logity/internal/domain/usecase/operating"
	"logity/internal/domain/usecase/room"
	"logity/internal/infrustructure/genHash/bcrypt"
	"logity/internal/infrustructure/realtime/logs"
	room3 "logity/internal/infrustructure/realtime/room"
	"logity/internal/infrustructure/repository/auth"
	log3 "logity/internal/infrustructure/repository/log"
	room2 "logity/internal/infrustructure/repository/room"
	"logity/internal/infrustructure/repository/users"
	"logity/internal/infrustructure/tokenManager"
	"logity/pkg/cetrifugo"
	"logity/pkg/neo4j"
	"logity/pkg/postgres"
	"net/http"
)

func Run(cfg *config.Config) {
	centrifugoClient, err := cetrifugo.NewCentrifugo(&cfg.Centrifugo)
	if err != nil {
		log.Fatalf("error centrifugo init client: %s", err)
	}

	fmt.Println(centrifugoClient.Info(context.Background())) //todo убрать

	dbClient, err := postgres.New(&cfg.Database)
	if err != nil {
		log.Fatalf("error init client db: %s", err)
	}

	neo4jDriver, err := neo4j.NewDriverNeo4j(&cfg.Neo4j)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error liquibase driver init: %s", err))
	}

	roomPublisher := room3.NewPublisher(centrifugoClient)
	logPublisher := logs.NewPublisher(centrifugoClient)

	generator := bcrypt.NewGenerator(&cfg.Auth)

	authRepo := auth.NewUserRepository(dbClient, generator)

	roomRepo := room2.NewRoomRepository(neo4jDriver, &cfg.Neo4j)

	userRepo := users.NewRepository(neo4jDriver, &cfg.Neo4j)

	logRepo := log3.NewRepository(neo4jDriver, &cfg.Neo4j)

	tokenMng := tokenManager.NewTokenManager(cfg)

	authUc := authUsecase.NewUserUsecase(authRepo, userRepo, tokenMng, cfg)
	roomUc := room.NewRoomUsecase(roomRepo, roomPublisher)
	logUc := log2.NewUsecase(logRepo, logPublisher)
	operatingUc := operating.NewUsecase(roomUc, authUc)

	env := usecase.NewEnv(roomUc, authUc, operatingUc, logUc)

	r := rest.NewRouter()
	rest.RegisterRouting(r, env, &cfg.App)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ApiPort), r); err != nil {
		log.Fatalf("error http server %s", err)
	}
}
