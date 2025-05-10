# 🛒 Resp

An opinionated gin response helpers.

## Installation

```bash
go get github.com/tsingshaner/gin/resp
```

## Usage

```go
package main
```

## SSE

```go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/tsingshaner/gin/resp/sse"
)

func main() {
  r := gin.Default()
  r.GET("/events", func(c *gin.Context) {
    if err := sse.Upgrader(c); err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
      return
    }

    body := make(chan []byte)

    go func() {
      for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        body <- []byte("hello")
      }
      close(body)
    }()

    c.Stream(func(w io.Writer) bool {
      c.SSEvent("open", "success")
      c.Writer.Flush()

      select {
      case <-c.Request.Context().Done():
        // logger.Debug("client close")
      case data := <-body:
        c.SSEvent("message", string(data))
      }
      return false
    })
  })

  r.Run(":8080")
}
```
