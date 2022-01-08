package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	Code string `json:"code" xml:"code"`
	Desc string `json:"desc" xml:"desc"`
}

func main() {

	e := echo.New()

	e.POST("/:parentId/:childId", func(c echo.Context) error {

		var (
			requestBody RequestBody
			parentId    string
			childId     string
			version     string
			types       string
		)

		// bind param to variable
		parentId = c.Param("parentId")
		childId = c.Param("childId")

		// bind param to variable
		version = c.QueryParam("version")
		types = c.QueryParam("types")

		// bind request to struct
		if err := c.Bind(&requestBody); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		fmt.Printf("%+v\n", parentId)
		fmt.Printf("%+v\n", childId)
		fmt.Printf("%+v\n", version)
		fmt.Printf("%+v\n", types)
		fmt.Printf("%+v\n", requestBody)

		return c.String(http.StatusOK, "hello world")

	})

	e.Logger.Fatal(e.Start(":8080"))

}
