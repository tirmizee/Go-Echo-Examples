package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ServerAttribute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("RequestId")
		c.Set("RequestId", uuid.New().String())
		return next(c)
	}
}

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main() {

	e := echo.New()

	e.GET("/", Handler, ServerAttribute)
	e.GET("/user", Handler)

	e.Logger.Fatal(e.Start(":8080"))

}
