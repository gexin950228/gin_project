package chapter04

import "github.com/gin-gonic/gin"

func Router(chapt04 *gin.RouterGroup) {
	chapt04.GET("/bind1", ShouldBind1)
	chapt04.POST("/doBind1", DoShouldBind1)
	chapt04.GET("/bind2", ShouldBind2)
	chapt04.POST("/doBind2", DoShouldBind2)
	chapt04.GET("/to_valid1", ToValidData)
	chapt04.POST("/do_valid1", DoValidData)

	chapt04.GET("/bind_uri/:id/:name/:addr/:age", BindUri)

}
