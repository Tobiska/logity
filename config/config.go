package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type (
	Config struct {
		Database
		Auth
		App
	}

	App struct {
		Host string
	}

	Auth struct {
		PasswordHashCost     int    `env:"PASSWORD_HASH_COST" env-default:"5"`
		SecretAccessKey      string `env:"SECRET_ACCESS_KEY"`
		SecretRefreshKey     string `env:"SECRET_REFRESH_KEY"`
		AccessTokenTTLInSec  int    `env:"ACCESS_TOKEN_TTL_IN_SEC" env-default:"350"`
		RefreshTokenTTLInSec int    `env:"ACCESS_TOKEN_TTL_IN_SEC" env-default:"3600"`
	}

	Database struct {
		Dsn             string `env:"POSTGRES_DSN"`
		MaxIdleConn     int    `env:"POSTGRES_MAX_IDLE_CONN" env-default:"3"`
		MaxLifeTimeConn int    `env:"POSTGRES_LIFETIME_CONN" env-default:"3"`
	}
)

var configInstance *Config
var configErr error
var once sync.Once

func ReadConfig() (*Config, error) {
	once.Do(func() {
		configErr = cleanenv.ReadEnv(&configInstance)
	})
	return configInstance, configErr
}
