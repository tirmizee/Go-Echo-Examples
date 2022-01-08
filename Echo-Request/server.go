package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/users/:name", func(c echo.Context) error {

		name := c.Param("name")

		code := c.QueryParam("code")

		types := c.FormValue("type")

		return c.String(http.StatusOK, name+code+types)

	})

	e.Logger.Fatal(e.Start(":8080"))

}
