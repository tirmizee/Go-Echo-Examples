package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main() {

	e := echo.New()

	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))

	e.GET("/new", Handler)
	e.GET("/get", Handler)
	e.GET("/put", Handler)
	e.GET("/public/javascripts/1", Handler)
	e.GET("/public/javascripts/2", Handler)
	e.GET("/user/1/order/1", Handler)
	e.GET("/user/2/order/2", Handler)

	e.Logger.Fatal(e.Start(":8080"))
}
