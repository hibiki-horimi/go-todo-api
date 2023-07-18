package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
	DB     DB
	OS     OS
}

type Server struct {
	Port    int    `envconfig:"Server_PORT" default:"8080"`
	Host    string `envconfig:"SERVER_HOST" default:"127.0.0.1"`
	Version string `envconfig:"VERSION" default:"default"`
}

type DB struct {
	Driver   string `envconfig:"DB_DRIVER" default:"postgres"`
	User     string `envconfig:"DB_USER_NAME" default:"postgres"`
	Password string `envconfig:"DB_USER_PASSWORD" default:"postgres"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"develop"`
	Schema   string `envconfig:"DB_SCHEMA" default:"public"`
}

type OS struct {
	TZ    string `envconfig:"TZ" default:"Asia/Tokyo"`
	Env   string `envconfig:"ENV" default:"local"`
	Debug string `envconfig:"DEBUG" default:"false"`
}

func New() (*Config, error) {
	var server Server
	if err := envconfig.Process("SERVER", &server); err != nil {
		return nil, fmt.Errorf("failed too get server env, err: %s", err)
	}
	var db DB
	if err := envconfig.Process("DB", &db); err != nil {
		return nil, fmt.Errorf("failed too get DB env, err: %s", err)
	}
	var os OS
	if err := envconfig.Process("OS", &os); err != nil {
		return nil, fmt.Errorf("failed too get OS env, err: %s", err)
	}
	return &Config{
		Server: server,
		DB:     db,
		OS:     os,
	}, nil
}
