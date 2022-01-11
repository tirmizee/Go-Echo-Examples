package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")
		return next(c)
	}
}

func ServerAttribute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("RequestId", uuid.New().String())
		return next(c)
	}
}

func Handler(c echo.Context) error {
	requestId := c.Get("RequestId").(string)
	return c.String(http.StatusOK, requestId)
}

func main() {

	e := echo.New()

	// add middlewares
	e.Use(ServerHeader, ServerAttribute)

	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(":8080"))

}
