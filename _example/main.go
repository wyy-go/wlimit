package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wlimit"
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
