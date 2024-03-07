package chapter11

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiAxios(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}
