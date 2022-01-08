- go mod init echo-request
- go get github.com/labstack/echo/v4

### Demo

```go

func main() {

	e := echo.New()

	e.POST("/users/:name", func(c echo.Context) error {

		name := c.Param("name")

		code := c.QueryParam("code")

		types := c.FormValue("type")

		return c.String(http.StatusOK, name+code+types)

	})

	e.Logger.Fatal(e.Start(":8080"))

}

```

        curl -XPOST -d 'type=A' 'http://localhost:8080/users/pratya?code=001'
