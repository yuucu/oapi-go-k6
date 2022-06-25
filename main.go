package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuucu/oapi-go-k6/generated/user_handler"

	emiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	handler := user_handler.NewGeneratedHandler()

	e := echo.New()
	e.Use(
		emiddleware.Recover(),
		emiddleware.CORSWithConfig(emiddleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}),
	)
	user_handler.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 3000)))
}
