package request

import (
	"net/http"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
)

type CustomBinder struct {
	binder echo.DefaultBinder
}

func InitBinder() echo.Binder {
	return &CustomBinder{
		binder: echo.DefaultBinder{},
	}
}

func (cb *CustomBinder) Bind(v any, c echo.Context) error {
	if err := cb.binder.Bind(v, c); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func InitValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(v any) error {
	if err := cv.validator.Struct(v); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}
