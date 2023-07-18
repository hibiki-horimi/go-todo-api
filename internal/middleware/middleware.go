package middleware

import (
	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Setup(e *echo.Echo, gdb *gorm.DB, rdb *postgres.Postgres) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}", "uri":"${uri}", "status":"${status}", "error":"${error}", "latency":"${latency}"}` + "\n",
	}))
	e.Use(setDB(gdb))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			return err
		},
	}))
}

func setDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetRequest(c.Request().WithContext(postgres.SetDB(c.Request().Context(), db)))
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
