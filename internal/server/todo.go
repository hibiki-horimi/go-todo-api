package server

import (
	"net/http"

	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
	"github.com/hibiki-horimi/go-todo-api/internal/domain"
	"github.com/hibiki-horimi/go-todo-api/internal/server/request"
	"github.com/hibiki-horimi/go-todo-api/internal/server/response"
	echo "github.com/labstack/echo/v4"
)

type Todo interface {
	List(c echo.Context) error
	Create(c echo.Context) error
}

type todo struct {
	rdb *postgres.Postgres
}

func (s *todo) List(c echo.Context) error {
	ctx := c.Request().Context()

	todoList, err := s.rdb.Todo.Find(ctx, &domain.Todo{})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.ToTodoList(todoList))
}

func (s *todo) Create(c echo.Context) error {
	var req request.CreateTodo
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	todo := req.ToTodo()

	if err := s.rdb.Todo.Create(ctx, todo); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.ToTodo(todo))
}
