- go mod init middleware
- go get github.com/labstack/echo/v4
- go get github.com/labstack/echo/v4/middleware
- go get github.com/google/uuid

### Demo

```go

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")
		return next(c)
	}
}

func ServerAttribute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("RequestId", uuid.New().String())
		return next(c)
	}
}

func Handler(c echo.Context) error {
	requestId := c.Get("RequestId").(string)
	return c.String(http.StatusOK, requestId)
}

func main() {

	e := echo.New()

	// add middlewares
	e.Use(ServerHeader, ServerAttribute)

	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(":8080"))

}

```

        http://localhost:8080/

        42826a34-17e4-488f-bc99-c1cac850f8fe