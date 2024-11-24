package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Env string `env:"ENV" env-required:"true"`
	App
	HTTP
	PG
}

type App struct {
	Name    string `env-required:"true" env:"APP_NAME"`
	Version string `env-required:"true" env:"APP_VERSION"`
}

type HTTP struct {
	Host string `env-required:"true" env:"HTTP_HOST"`
	Port string `env-required:"true" env:"HTTP_PORT"`
}

type PG struct {
	PoolSize int    `env-required:"true" env:"PG_POOL_SIZE"`
	URL      string `env-required:"true" env:"PG_URL"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Error update ENV config: %v", err)
	}

	return cfg
}
