- go mod init requestid
- go get github.com/labstack/echo/v4
- go get github.com/labstack/echo/v4/middleware

### Demo

```go

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main() {

	e := echo.New()

	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))

	e.GET("/new", Handler)
	e.GET("/get", Handler)
	e.GET("/put", Handler)
	e.GET("/public/javascripts/1", Handler)
	e.GET("/public/javascripts/2", Handler)
	e.GET("/user/1/order/1", Handler)
	e.GET("/user/2/order/2", Handler)

	e.Logger.Fatal(e.Start(":8080"))
}

```
		http://localhost:8080/ne  					// 404
		http://localhost:8080/new 					// 200
		http://localhost:8080/old 					// 200
		http://localhost:8080/api/get 				// 200
		http://localhost:8080/get					// 200
		http://localhost:8080/api/put 				// 200
		http://localhost:8080/put					// 200
		http://localhost:8080/public/javascripts/1	// 200
		http://localhost:8080/js/1 					// 200
		http://localhost:8080/public/javascripts/2	// 200
		http://localhost:8080/js/2 					// 200
		http://localhost:8080/users/1/orders/1 		// 200
		http://localhost:8080/user/1/order/1 		// 200
		http://localhost:8080/users/2/orders/2 		// 200
		http://localhost:8080/user/2/order/2 		// 200