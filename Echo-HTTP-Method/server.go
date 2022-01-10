package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler

func HandlerGET(c echo.Context) error {
	return c.String(http.StatusOK, "GET Method")
}

func HandlerPOST(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func HandlerPUT(c echo.Context) error {
	return c.String(http.StatusOK, "PUT Method")
}

func HandlerDELETE(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE Method")
}

func main() {

	e := echo.New()

	e.GET("/GET", HandlerGET)
	e.PUT("/PUT", HandlerPUT)
	e.POST("/POST", HandlerPOST)
	e.DELETE("/DELETE", HandlerDELETE)

	e.Logger.Fatal(e.Start(":8080"))

}
