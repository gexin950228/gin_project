package chapter03

import "github.com/gin-gonic/gin"

func Router(chapt03 *gin.RouterGroup) {
	chapt03.GET("/tpl1", Tpl1)
	chapt03.GET("/func1", Func1)
}
