package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	ID   string `param:"id"`
	Name string `param:"name"`
}

func main() {

	// Echo instance
	e := echo.New()

	e.GET("/:id/:name", func(c echo.Context) error {

		var productRequest ProductRequest

		if err := c.Bind(&productRequest); err != nil {
			return c.String(http.StatusBadRequest, "Can not bind data")
		}

		return c.String(http.StatusOK, productRequest.ID+","+productRequest.Name)

	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
