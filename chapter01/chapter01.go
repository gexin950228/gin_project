package chapter01

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter01/index.html", nil)
}

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello...")
}
