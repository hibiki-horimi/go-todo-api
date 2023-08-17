package route

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/hibiki-horimi/go-todo-api/internal/server"
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

		todo := api.Group("/todos")
		{
			todo.GET("", srv.Todo.List)
			todo.POST("", srv.Todo.Create)

			task := todo.Group("/:id")
			{
				task.GET("", srv.Todo.Get)
				task.PUT("", srv.Todo.Update)
				task.DELETE("", srv.Todo.Delete)
			}
		}
	}
}
