package user_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GeneratedHandler struct{}

func NewGeneratedHandler() *GeneratedHandler {
	return &GeneratedHandler{}
}

// (GET /user/{id})
func (h GeneratedHandler) GetUser(c echo.Context, id int) error {
	return c.String(http.StatusOK, "ok")
}

// (POST /user/{id})
func (h GeneratedHandler) PostUser(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// (DELETE /user/{id})
func (h GeneratedHandler) DeleteUser(c echo.Context, id int) error {
	return c.String(http.StatusOK, "ok")
}
