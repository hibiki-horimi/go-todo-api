package server

import (
	"errors"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
	"github.com/hibiki-horimi/go-todo-api/internal/domain"
	"github.com/hibiki-horimi/go-todo-api/internal/server/request"
	"github.com/hibiki-horimi/go-todo-api/internal/server/response"
)

type Todo interface {
	Get(c echo.Context) error
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type todo struct {
	rdb *postgres.Postgres
}

func (s *todo) Get(c echo.Context) error {
	var req request.GetTodo
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	res, err := s.rdb.Todo.Get(ctx, req.ToTodo())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return err
	}
	return c.JSON(http.StatusOK, response.ToTodo(res))
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

	return c.JSON(http.StatusCreated, response.ToTodo(todo))
}

func (s *todo) Update(c echo.Context) error {
	var req request.UpdateTodo
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	todo := req.ToTodo()

	if err := s.rdb.Todo.Update(ctx, todo); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.ToTodo(todo))
}

func (s *todo) Delete(c echo.Context) error {
	var req request.DeleteTodo
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	_, err := s.rdb.Todo.Get(ctx, req.ToTodo())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return err
	}

	if err := s.rdb.Todo.Delete(ctx, req.ToTodo()); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
