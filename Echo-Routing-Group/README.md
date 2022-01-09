- go mod init echo-group
- go get -u github.com/labstack/echo/v4

### Demo

```go

// Handler

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func Handler1(c echo.Context) error {
	return c.String(http.StatusOK, "Admin group")
}

func Handler2(c echo.Context) error {
	return c.String(http.StatusOK, "Admin group")
}

func Handler3(c echo.Context) error {
	return c.String(http.StatusOK, "Admin group")
}

func main() {

	e := echo.New()

	// no group
	e.GET("/users", Handler)

	// admin group
	g := e.Group("/admin")
	g.GET("/users", Handler1)
	g.GET("/users/:name", Handler2)
	g.GET("/files/*", Handler3)

	e.Logger.Fatal(e.Start(":8080"))

}

```

        http://localhost:8080/users
        http://localhost:8080/admin/users
        http://localhost:8080/admin/users/tir
        http://localhost:8080/admin/files/1