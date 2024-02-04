package chapter07

import (
	"github.com/gin-gonic/gin"
)

func Router(chap07 *gin.RouterGroup) {
	chap07.GET("/gorm_test", GormTest)
}
