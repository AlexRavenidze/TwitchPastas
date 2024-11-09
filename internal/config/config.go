package config

import (
	"time"
)

type Configuration struct {
	AppConfig appConfig
	Postgres  postgresConfig
}

type appConfig struct {
	AppName      string        `envconfig:"APP_NAME" required:"true"`
	Port         string        `envconfig:"PORT" required:"true"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"60s"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"60s"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" default:"5s"`
	BodyLimit    int           `envconfig:"BODY_LIMIT" default:"4194304"`
	LogLevel     string        `envconfig:"LOG_LEVEL" default:"info"`
}

type postgresConfig struct {
}
