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

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.binder.Bind(i, c); err != nil {
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

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}
