package app

import (
	"log"
	"logity/config"
	"logity/pkg/postgres"
)

func Run(cfg *config.Config) {
	dbClient, err := postgres.New(&cfg.Database)
	if err != nil {
		log.Fatalf("error init client db: %s", err)
	}

}
