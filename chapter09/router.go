package chapter09

import "github.com/gin-gonic/gin"

func Router(chapt09 *gin.RouterGroup) {
	chapt09.GET("/test_gorm", SessionTest)
}
