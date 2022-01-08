- go mod init bind-param
- go get github.com/labstack/echo/v4

### Demo

```go

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

```


        http://localhost:1323/12321/kfsdk