# wlimit


![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wlimit?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wlimit)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wlimit)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wlimit/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wlimit)](https://goreportcard.com/report/github.com/wyy-go/wlimit)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wlimit?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wlimit/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wlimit)


Limit size of POST requests for Gin framework

## Example

```go
package main

import (
  "net/http"

  "github.com/wyy-go/wlimit"
  "github.com/gin-gonic/gin"
  "log"
)

func handler(ctx *gin.Context) {
  val := ctx.PostForm("b")
  if len(ctx.Errors) > 0 {
    return
  }
  ctx.String(http.StatusOK, "got %s\n", val)
}

func main() {
  r := gin.Default()
  r.Use(wlimit.New(wlimit.WithLimit(10)))
  r.POST("/", handler)
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
