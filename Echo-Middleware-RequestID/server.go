package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, c.Response().Header().Get(echo.HeaderXRequestID))
}

func main() {

	e := echo.New()

	// comfig with default
	e.Use(middleware.RequestID())

	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(":8080"))
}
