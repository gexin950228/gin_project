package chapter11

import "github.com/gin-gonic/gin"

func Router(chapt11 *gin.RouterGroup) {
	chapt11.GET("/apiAxios", ApiAxios)
	chapt11.GET("/books", GetBooks)
}
