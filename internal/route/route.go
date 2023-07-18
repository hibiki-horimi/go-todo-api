package route

import (
	"net/http"

	"github.com/hibiki-horimi/go-todo-api/internal/server"
	echo "github.com/labstack/echo/v4"
)

type router struct {
	e      *echo.Echo
	server *server.Server
}

func New(e *echo.Echo, srv *server.Server) {
	r := &router{e: e, server: srv}
	api := r.e.Group("/api")
	{
		api.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "healthy")
		})

		task := api.Group("/tasks")
		{
			task.GET("/", srv.Todo.List)
			task.POST("/", srv.Todo.Create)
		}
	}
}
