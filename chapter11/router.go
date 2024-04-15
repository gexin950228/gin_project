package chapter11

import "github.com/gin-gonic/gin"

func Router(chapt11 *gin.RouterGroup) {
	chapt11.GET("/apiAxios", ApiAxios)
	chapt11.GET("/books", GetBooks)
	chapt11.GET("/book_detail", GetBookDetail)
	chapt11.POST("/postTest", PostTest)
	chapt11.POST("/fileUpload", FileUpload)
	chapt11.POST("/filesUpload", FilesUpload)
}
