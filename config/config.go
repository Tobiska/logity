package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type (
	Config struct {
		App
		Auth
		Database
		Neo4j
		Centrifugo
	}

	Neo4j struct {
		Host     string `env:"NEO4J_HOST"`
		Database string `env:"NEO4J_DATABASE"`
		Username string `env:"NEO4J_USERNAME"`
		Password string `env:"NEO4J_PASSWORD"`
	}

	Centrifugo struct {
		ApiHost       string `env:"CENTRIFUGO_API_HOST"`
		ClientHost    string `env:"CENTRIFUGO_CLIENT_HOST"`
		ApiKey        string `env:"CENTRIFUGO_API_KEY"`
		SecretKey     string `env:"CENTRIFUGO_SECRET_KEY"`
		TokenTTLInSec int    `env:"CENTRIFUGO_TOKEN_TTL_IN_SEC" env-default:"10"`
	}

	App struct {
		Host    string `env:"HOST_NAME" env-default:"LOGITY"`
		ApiPort string `env:"API_PORT" env-default:":8080"`
	}

	Auth struct {
		PasswordHashCost     int    `env:"PASSWORD_HASH_COST" env-default:"5"`
		SecretAccessKey      string `env:"SECRET_ACCESS_KEY"`
		SecretRefreshKey     string `env:"SECRET_REFRESH_KEY"`
		AccessTokenTTLInSec  int    `env:"ACCESS_TOKEN_TTL_IN_SEC" env-default:"350"`
		RefreshTokenTTLInSec int    `env:"REFRESH_TOKEN_TTL_IN_SEC" env-default:"3600"`
	}

	Database struct {
		Dsn             string `env:"POSTGRES_DSN"`
		MaxIdleConn     int    `env:"POSTGRES_MAX_IDLE_CONN" env-default:"3"`
		MaxLifeTimeConn int    `env:"POSTGRES_LIFETIME_CONN" env-default:"3"`
	}
)

var configInstance *Config
var configErr error

func ReadConfig() (*Config, error) {
	if configInstance == nil {
		var readConfigOnce sync.Once
		readConfigOnce.Do(func() {
			configInstance = &Config{}
			configErr = cleanenv.ReadEnv(configInstance)
		})
	}

	return configInstance, configErr
}
