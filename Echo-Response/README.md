- go mod init echo-response
- go get github.com/labstack/echo/v4

### Demo

```go

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

```

        http://localhost:8080/string
        http://localhost:8080/json
        http://localhost:8080/json_pretty
        http://localhost:8080/xml
        http://localhost:8080/html