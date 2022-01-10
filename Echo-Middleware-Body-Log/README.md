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