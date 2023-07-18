package server

import (
	"github.com/hibiki-horimi/go-todo-api/internal/config"
	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
)

type Server struct {
	Todo Todo
}

func New(rdb *postgres.Postgres, conf *config.Config) *Server {
	return &Server{
		Todo: &todo{rdb: rdb},
	}
}
