- go mod init auth-key
- go get github.com/labstack/echo/v4
- go get github.com/labstack/echo/v4/middleware

### Demo

```go

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main() {

	e := echo.New()

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			fmt.Println(key)
			return key == "valid-key", nil
		},
	}))

	e.GET("/auth", Handler)

	e.Logger.Fatal(e.Start(":8080"))

}

```


		curl -XGET -H 'api-key: valid-key' 'http://localhost:8080/auth'