- go mod init echo-http
- go get github.com/labstack/echo/v4 

### Demo

```go

// Handler

func HandlerGET(c echo.Context) error {
	return c.String(http.StatusOK, "GET Method")
}

func HandlerPOST(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func HandlerPUT(c echo.Context) error {
	return c.String(http.StatusOK, "PUT Method")
}

func HandlerDELETE(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE Method")
}

func main() {

	e := echo.New()

	e.GET("/GET", HandlerGET)
	e.PUT("/PUT", HandlerPUT)
	e.GET("/POST", HandlerPOST)
	e.GET("/DELETE", HandlerDELETE)

	e.Logger.Fatal(e.Start(":8080"))

}

```


