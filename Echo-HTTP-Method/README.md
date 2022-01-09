- go mod init echo-http
- go get github.com/labstack/echo/v4 

### Http Method

```go

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (e *Echo) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodDelete, path, h, m...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (e *Echo) GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodGet, path, h, m...)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Echo) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodHead, path, h, m...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Echo) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodOptions, path, h, m...)
}

// PATCH registers a new PATCH route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Echo) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodPatch, path, h, m...)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Echo) POST(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodPost, path, h, m...)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Echo) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodPut, path, h, m...)
}

```

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


