package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler

func Handler1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Handler2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Handler3(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {

	e := echo.New()

	// Route
	e.GET("/users", Handler1)
	e.GET("/users/:name", Handler2)
	e.GET("/files/*", Handler3)

	e.Logger.Fatal(e.Start(":8080"))

}
