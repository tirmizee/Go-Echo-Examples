- go get github.com/labstack/echo
- go get github.com/labstack/echo/v4


### demo

```go

type ProductRequest struct {
	ID   string `query:"id"`
	Name string `query:"name"`
}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		var productRequest ProductRequest

		if err := c.Bind(&productRequest); err != nil {
			return c.String(http.StatusBadRequest, "Can not bind data")
		}

		return c.String(http.StatusOK, productRequest.ID+","+productRequest.Name)

	})
	e.Logger.Fatal(e.Start(":1323"))
}

```

        http://localhost:1323/?id=2222&name=fafaa

        