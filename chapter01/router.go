package chapter01

import "github.com/gin-gonic/gin"

func Router(chapt01 *gin.RouterGroup) {
	chapt01.GET("/hello", Hello)
	chapt01.GET("/user", User)
}
