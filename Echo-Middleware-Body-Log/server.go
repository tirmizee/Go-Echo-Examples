package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Println("req:", string(reqBody))
	fmt.Println("res:", string(resBody))
}

func HandlerPOST(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func main() {

	e := echo.New()

	// manual config
	e.Use(middleware.BodyDump(BodyDumpHandler))

	// custom config
	// e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))

	// rounter
	e.POST("/", HandlerPOST)

	e.Logger.Fatal(e.Start(":8080"))

}
