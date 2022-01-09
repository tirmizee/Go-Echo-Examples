- go mod init mid-cors
- go get github.com/labstack/echo/v4

### Defualt config

```go 

DefaultCORSConfig = CORSConfig{
  Skipper:      DefaultSkipper,
  AllowOrigins: []string{"*"},
  AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
}

```

### Demo

```go

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "cors")
}

func main() {

	e := echo.New()

	// default config
	// e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// custom config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", Handler)

	e.Logger.Fatal(e.Start(":8080"))

}

```