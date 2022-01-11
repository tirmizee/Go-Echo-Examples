- go mod init error-handler
- go get github.com/labstack/echo/v4

### Demo

```go

// Http error response model
type HTTPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// new http error
func NewHTTPError(code string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: msg,
	}
}

// Error method for implement error interface
func (e *HTTPError) Error() string {
	return e.Code + ": " + e.Message
}

func HTTPErrorHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		if httpError, ok := err.(*HTTPError); ok {
			c.Logger().Error(c.JSON(http.StatusOK, httpError))
		} else {
			e.DefaultHTTPErrorHandler(err, c)
		}

	}
}

func GetUser(id string) (interface{}, error) {
	return nil, NewHTTPError("ERR001", "user not found")
}

func HandlerPOST(c echo.Context) error {

	id := c.QueryParam("id")

	_, err := GetUser(id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "hello world")
}

func main() {

	e := echo.New()

	e.HTTPErrorHandler = HTTPErrorHandler(e)

	e.POST("/", HandlerPOST)

	e.Logger.Fatal(e.Start(":8080"))
}


```

        curl -XPOST -H "Content-type: application/json" 'http://localhost:8080/?id=1'

        {"code":"ERR001","message":"user not found"}