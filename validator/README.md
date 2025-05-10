# 🛂 Validator

Just a gin requuest params validator middleware.

## Installation

```bash
go get github.com/tsingshaner/gin/validator
```

## Usage

```go
package main

type HandlerBody struct {
    Name string `json:"name" binding:"required"`
}

type HandlerQuery struct {
    Age int `form:"age" binding:"required"`
}

type HandlerParams struct {
    ID int `uri:"id" binding:"required"`
}

type HandlerHeader struct {
    Token string `header:"token" binding:"required"`
}

func main() {
    r := gin.Default()
    r.Use(v.Header[]())
    r.POST("/api/users/:id", v.Params[](), v.Query[](), v.Body[](), Handler)
    r.Run()
}

func Handler(c *gin.Context) {
    header := v.GetHeader[](c)
    params := v.GetParams[](c)
    query := v.GetQuery[](c)
    body := v.GetBody[](c)

    // do something
}
```
