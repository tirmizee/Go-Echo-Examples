package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func HandlerPOST(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func main() {

	e := echo.New()

	// config
	e.Use(middleware.BodyLimit("1K"))

	// contom config
	// e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{}))

	// router
	e.POST("/", HandlerPOST)

	e.Logger.Fatal(e.Start(":8080"))

}
