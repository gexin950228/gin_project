package chapter05

import "github.com/gin-gonic/gin"

func Router(chapt05 *gin.RouterGroup) {
	chapt05.GET("/auth", gin.BasicAuth(gin.Accounts{
		"zs": "123456",
		"ls": "123456",
		"ww": "123456",
	}), gin.WrapF(WrapTest), AuthTest)
}
