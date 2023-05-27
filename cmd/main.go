package main

import (
	"log"
	"logity/config"
	"logity/internal/app"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("read config err: %s\n", err)
	}

	app.Run(cfg)
}
