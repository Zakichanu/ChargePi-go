package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

type UI struct {
}

func NewUi() *UI {
	return &UI{}
}

func (u *UI) Serve(url string) {
	r := gin.Default()

	r.Use(spa.Middleware("/", "./ui/build"))

	err := r.Run(url)
	if err != nil {
		return
	}
}