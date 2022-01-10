Body dump middleware captures the request and response payload 

- go mod init mid-body-log
- github.com/labstack/echo/middleware
- go get github.com/labstack/echo/v4

### Configuration

```go

BodyDumpConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Handler receives request and response payload.
  Handler BodyDumpHandler

}

```

### Demo

```go

func BodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Println("req:", string(reqBody))
	fmt.Println("res:", string(resBody))
}

func HandlerPOST(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func main() {

	e := echo.New()

	// manual config
	e.Use(middleware.BodyDump(BodyDumpHandler))

	// custom config
	// e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))

	// router
	e.POST("/", HandlerPOST)

	e.Logger.Fatal(e.Start(":8080"))

}

```

    curl -XPOST -H "Content-type: application/json" -d '{"key1":true, "key2":"text","key3":123}' 'http://localhost:8080/'

    req: {"key1":true, "key2":"text","key3":123}
    res: POST Method