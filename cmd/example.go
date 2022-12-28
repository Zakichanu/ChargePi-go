package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

func main() {
	r := gin.Default()

	r.GET("/api/test", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	r.Use(spa.Middleware("/", "./ui/build"))

	err := r.Run("0.0.0.0:8001")
	if err != nil {
		return
	}
}
