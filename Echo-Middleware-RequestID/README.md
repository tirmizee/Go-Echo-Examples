- go mod init requestid
- go get github.com/labstack/echo/v4
- go get github.com/labstack/echo/v4/middleware
- go get github.com/google/uuid

### Demo

```go

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, c.Response().Header().Get(echo.HeaderXRequestID))
}

func main() {

	e := echo.New()

	// comfig with default
	e.Use(middleware.RequestID())

	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(":8080"))
}

```

        http://localhost:8080/

		curl -XGET -H 'X-Request-ID: 123456' 'http://localhost:8080'