package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func skip(c echo.Context) bool {

	if c.Request().URL.Path == "/user" {
		return true
	}

	if c.RealIP() == "localhost" {
		return true
	}

	return false
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if skip(c) {
			return next(c)
		}

		fmt.Println("middleware processing")

		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")

		return next(c)
	}
}

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main() {
	e := echo.New()

	e.Use(ServerHeader)

	e.GET("/", Handler)
	e.GET("/user", Handler)

	e.Logger.Fatal(e.Start(":8080"))
}
