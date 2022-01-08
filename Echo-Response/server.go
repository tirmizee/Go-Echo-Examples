package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Transfer struct {
	Code string
	Desc string
}

func main() {

	e := echo.New()

	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Transfer{"001", "Success"})
	})

	e.GET("/json_pretty", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, &Transfer{"001", "Success"}, " ")
	})

	e.GET("/xml", func(c echo.Context) error {
		return c.XML(http.StatusOK, &Transfer{"001", "Success"})
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
