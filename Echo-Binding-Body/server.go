package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	ID   string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {

		var productRequest ProductRequest

		if err := c.Bind(&productRequest); err != nil {
			return c.String(http.StatusBadRequest, "Can not bind data")
		}

		return c.String(http.StatusOK, productRequest.ID+","+productRequest.Name)

	})
	e.Logger.Fatal(e.Start(":8080"))
}
