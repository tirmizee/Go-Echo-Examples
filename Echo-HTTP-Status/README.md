- go mod init echo-status
- go get github.com/labstack/echo/v4

### Demo

```go

// Handler

func Handler1(c echo.Context) error {
	return c.String(http.StatusOK, "StatusOK")
}

func Handler2(c echo.Context) error {
	return c.String(http.StatusForbidden, "StatusForbidden")
}

func Handler3(c echo.Context) error {
	return c.String(http.StatusBadRequest, "StatusBadRequest")
}

func Handler4(c echo.Context) error {
	return c.String(http.StatusFound, "StatusFound")
}

func Handler5(c echo.Context) error {
	return c.String(http.StatusGatewayTimeout, "StatusGatewayTimeout")
}

func Handler6(c echo.Context) error {
	return c.String(http.StatusBadGateway, "StatusBadGateway")
}

func main() {

	e := echo.New()

	e.GET("/StatusOK", Handler1)
	e.GET("/StatusForbidden", Handler2)
	e.GET("/StatusBadRequest", Handler3)
	e.GET("/StatusFound", Handler4)
	e.GET("/StatusGatewayTimeout", Handler5)
	e.GET("/StatusBadGateway", Handler6)

	e.Logger.Fatal(e.Start(":8080"))

}

```

        http://localhost:8080/StatusOK
        http://localhost:8080/StatusForbidden
        http://localhost:8080/StatusBadRequest
        http://localhost:8080/StatusFound
        http://localhost:8080/StatusGatewayTimeout
        http://localhost:8080/StatusBadGateway