# GoReply

A lightweight HTTP response helper library that simplifies response handling across multiple Go web frameworks. Write once, run anywhere.

## Why Reply?

- **Framework agnostic** - Works seamlessly with Gin, Echo, Fiber, and net/http
- **Consistent response format** - Standardized response structure across your API
- **Chainable API** - Clean and readable method chaining
- **Multiple formats** - JSON, XML, HTML, Binary, Text, and Streaming support
- **Smart error handling** - Code aliases and custom transformers

## Installation

```bash
go get github.com/chesta132/goreply
```

For specific framework adapters:

```bash
# Gin
go get github.com/gin-gonic/gin github.com/chesta132/goreply/adapter/gin

# Echo
go get github.com/labstack/echo/v4 github.com/chesta132/goreply/adapter/echo

# Fiber
go get github.com/gofiber/fiber/v2 github.com/chesta132/goreply/adapter/fiber
```

## Quick Start

### Setup Client

```go
import "github.com/chesta132/goreply/reply"

var client = reply.NewClient(reply.Client{
    CodeAliases: map[string]int{
        "NOT_FOUND":    404,
        "BAD_REQUEST":  400,
        "SERVER_ERROR": 500,
    },
    DefaultHeaders: map[string]string{
        "Content-Type": "application/json",
    },
    PaginationType: "page" // default: offset
})
```

### Basic Usage (net/http)

```go
import "github.com/chesta132/goreply/adapter/nethttp"

func handler(w http.ResponseWriter, r *http.Request) {
    rp := client.New(adapter.AdaptHttp(w, r))

    // Success response
    rp.Success(map[string]string{
        "message": "Hello World!",
    }).OkJSON()
}
```

**Output:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "message": "Hello World!"
  }
}
```

### With Gin

```go
import "github.com/chesta132/goreply/adapter/gin"

func handler(c *gin.Context) {
    rp := client.New(adapter.AdaptGin(c))

    users := []User{
        {ID: 1, Name: "Alice"},
        {ID: 2, Name: "Bob"},
    }

    rp.Success(users).OkJSON()
}
```

### With Echo

```go
import "github.com/chesta132/goreply/adapter/echo"

func handler(c echo.Context) error {
    rp := client.New(adapter.AdaptEcho(c))

    rp.Success("Hello from Echo!").OkText()
    return nil
}
```

### With Fiber

```go
import "github.com/chesta132/goreply/adapter/fiber"

func handler(c *fiber.Ctx) error {
    rp := client.New(adapter.AdaptFiber(c))

    rp.Success("<h1>Hello Fiber!</h1>").OkHTML()
    return nil
}
```

## Features

### Error Response

```go
// Simple error
rp.Error("NOT_FOUND", "User not found").FailJSON()

// Error with details
rp.Error("VALIDATION_ERROR", "Invalid input", reply.OptErrorPayload{
    Details: "Email format is invalid",
    Fields:   []string{"email"},
}).FailJSON(400)
```

**Output:**

```json
{
  "meta": {
    "status": "ERROR"
  },
  "data": {
    "code": "NOT_FOUND",
    "message": "User not found",
    "details": "Email format is invalid",
    "field": "email"
  }
}
```

### Pagination Total Based

```go
users := fetchUsers(limit, current)
total := getTotalUsers()

rp.Success(users).
    PaginateTotal(limit, current, total).
    OkJSON()
```

**Output:**

```json
{
  "meta": {
    "status": "SUCCESS",
    "pagination": {
      "next": 20,
      "hasNext": true,
      "current": 0,
      "total": 100
    }
  },
  "data": [...]
}
```

### Pagination Cursor Based

```go
users := fetchUsers(limit + 1, current)

rp.Success(users).
    PaginateCursor(limit, current).
    OkJSON()
```

**Output:**

```json
{
  "meta": {
    "status": "SUCCESS",
    "pagination": {
      "next": 20,
      "hasNext": true,
      "current": 0
    }
  },
  "data": [...]
}
```

### Cookie

```go
accessToken := createAccessToken()
refreshToken := createRefreshToken()

rp.SetCookies(
  http.Cookie{Name: "access_token", Value: accessToken},
  http.Cookie{Name: "refresh_token", Value: refreshToken},
)
```

### Multiple Deferred Functions

```go
rp.Defer(
    func() { cleanupTempFiles() },
    func() { closeConnection() },
    func() { log.Println("Request completed") },
)
```

### Response Formats

#### JSON

```go
rp.Success(data).OkJSON()           // 200
rp.Success(data).CreatedJSON()      // 201
rp.Error("ERR", "msg").FailJSON()   // from CodeAliases or 500
```

#### XML

```go
rp.Success(data).OkXML()            // 200
rp.Success(data).CreatedXML()       // 201
rp.Error("ERR", "msg").FailXML(400) // 400
```

#### Text

```go
rp.Success("Plain text").OkText()
rp.Success("Created!").CreatedText()
```

#### HTML

```go
rp.Success("<h1>Hello</h1>").OkHTML()
rp.Success("<p>Created</p>").CreatedHTML()
```

#### Binary

```go
imageData := []byte{...}
rp.Success(imageData).OkBinary()
```

#### Streaming

```go
file, _ := os.Open("video.mp4")
defer file.Close()

rp.Success(reply.Stream{
    Data:        file,
    ContentType: "video/mp4",
}).OkStream()
```

## Advanced Usage

### Custom Transformer

Transform the response structure before sending:

```go
client := reply.NewClient(reply.Client{
    Transformer: func(data any, meta reply.Meta) any {
        return map[string]any{
            "success": meta.Status == "SUCCESS",
            "payload": data,
            "timestamp": time.Now().Unix(),
        }
    },
})
```

### Finalizer Hook

Execute custom logic before sending responses:

```go
client := reply.NewClient(reply.Client{
    Finalizer: func(data any, meta reply.Meta) {
        // Log response before sending
        log.Printf("Sending response: status=%s", meta.Status)
    },
})
```

### Code Aliases

Map error codes to HTTP status codes:

```go
client := reply.NewClient(reply.Client{
    CodeAliases: map[string]int{
        "USER_NOT_FOUND":      404,
        "INVALID_CREDENTIALS": 401,
        "RATE_LIMITED":        429,
        "MAINTENANCE":         503,
    },
})

// Automatically uses status code from alias
rp.Error("RATE_LIMITED", "Too many requests").FailJSON()
// Response with status 429
```

### Default Headers

Set headers that will be applied to all responses:

```go
client := reply.NewClient(reply.Client{
    DefaultHeaders: map[string]string{
        "X-API-Version": "v1.0",
        "X-Powered-By":  "Reply-Go",
    },
})
```

## Response Structure

### Success Response

```json
{
  "meta": {
    "status": "SUCCESS",
    "information": "optional info",
    "pagination": {
      "Next": 20,
      "hasNext": true
    }
  },
  "data": {...}
}
```

### Error Response

```json
{
  "meta": {
    "status": "ERROR"
  },
  "data": {
    "code": "ERROR_CODE",
    "message": "Human readable message",
    "details": "Optional debug info",
    "field": "fieldName"
  }
}
```

## API Reference

### Core Methods

- `NewClient(config Client)` - Create a new client with configuration
- `Success(data any)` - Set success response
- `Error(code, message string, opt ...OptErrorPayload)` - Set error response
- `Info(information string)` - Set meta information
- `PaginateTotal(limit, offset, total int)` - Add pagination information with total based
- `PaginateCursor(limit, offset int)` - Add pagination information with cursor based
- `Defer(funcs ...func())` - Register functions to execute before sending response
- `SetCookies(cookies ...http.Cookie)` - Add Set-Cookie header by http.Cookie

### Response Methods

| Method                  | Status Code | Format | Returns |
| ----------------------- | ----------- | ------ | ------- |
| `NoContent()`           | 204         | -      | -       |
| `Redirect()`            | Custom      | -      | -       |
| `ReplyJSON()`           | Custom      | JSON   | error   |
| `OkJSON()`              | 200         | JSON   | error   |
| `CreatedJSON()`         | 201         | JSON   | error   |
| `FailJSON(code ...int)` | Custom/500  | JSON   | error   |
| `ReplyXML()`            | Custom      | XML    | error   |
| `OkXML()`               | 200         | XML    | error   |
| `CreatedXML()`          | 201         | XML    | error   |
| `FailXML(code ...int)`  | Custom/500  | XML    | error   |
| `ReplyText()`           | Custom      | Text   | error   |
| `OkText()`              | 200         | Text   | error   |
| `CreatedText()`         | 201         | Text   | error   |
| `ReplyHTML()`           | Custom      | HTML   | error   |
| `OkHTML()`              | 200         | HTML   | error   |
| `CreatedHTML()`         | 201         | HTML   | error   |
| `ReplyBinary()`         | Custom      | Binary | error   |
| `OkBinary()`            | 200         | Binary | error   |
| `CreatedBinary()`       | 201         | Binary | error   |
| `ReplyStream()`         | Custom      | Stream | error   |
| `OkStream()`            | 200         | Stream | error   |
| `CreatedStream()`       | 201         | Stream | error   |

## Framework Support

| Framework | Import Path       | Adapter Function  |
| --------- | ----------------- | ----------------- |
| net/http  | `adapter/nethttp` | `AdaptHttp(w, r)` |
| Gin       | `adapter/gin`     | `AdaptGin(c)`     |
| Echo      | `adapter/echo`    | `AdaptEcho(c)`    |
| Fiber     | `adapter/fiber`   | `AdaptFiber(c)`   |

## Real-World Examples

## User API with Pagination And Defer

```go
type FindManyPayload struct {
   Ids []string `json:"ids"`
}
func FindUsers(c *gin.Context) {
    rp := client.New(adapter.AdaptGin(c))

    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
    payload := FindManyPayload{}
    c.ShouldBindJSON(&payload)

    tx := db.Begin()

    // Defer cleanup - will run before response is sent
    rp.Defer(func() {
        if tx != nil {
            tx.Close()
        }
    })

    users, err := findUsers(tx, payload.Ids, limit+1, offset)
    if err != nil {
        rp.Error("FIND_FAILED", err.Error()).FailJSON()
        // tx is closed
        return
    }

    rp.Success(users).
        PaginateCursor(limit, offset).
        OkJSON()
}
```

### File Upload Handler

```go
func UploadFile(c echo.Context) error {
    rp := client.New(echoadapter.AdaptEcho(c))

    file, err := c.FormFile("file")
    if err != nil {
        rp.Error("INVALID_FILE", "No file uploaded").FailJSON(400)
        return nil
    }

    // Process file...

    rp.Success(map[string]any{
        "filename": file.Filename,
        "size":     file.Size,
    }).CreatedJSON()

    return nil
}
```

### Video Streaming

```go
func StreamVideo(w http.ResponseWriter, r *http.Request) {
    rp := client.New(nethttpadapter.AdaptHttp(w, r))

    file, err := os.Open("video.mp4")
    if err != nil {
        rp.Error("NOT_FOUND", "Video not found").FailJSON(404)
        return
    }
    defer file.Close()

    rp.Success(reply.Stream{
        Data:        file,
        ContentType: "video/mp4",
    }).OkStream()
}
```

## Contributing

Pull requests are welcome! Feel free to contribute by adding support for more frameworks or new features.

## License

MIT
