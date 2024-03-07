package chapter11

import "github.com/gin-gonic/gin"

func Router(chapt09 *gin.RouterGroup) {
	chapt09.GET("/apiAxios", ApiAxios)
}
