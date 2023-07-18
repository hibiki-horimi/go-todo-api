package request

import (
	"net/http"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func New() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}
