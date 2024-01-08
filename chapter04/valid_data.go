package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id      int    `form:"-" json:"id"`
	Title   string `form:"title" json:"title" binding:"omitempty,min=2"`
	Author  string `form:"author" json:"author"`
	Content string `form:"content" json:"content" binding:"min=4,max=10"`
	Gender  string `form:"gender" json:"gender" binding:"-"`
	Email   string `form:"email" binding:"email"`
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

func DoValidData(ctx *gin.Context) {
	var article Article
	errDataBind := ctx.ShouldBind(&article)
	if errDataBind != nil {
		fmt.Println(errDataBind.Error())
		ctx.String(http.StatusBadRequest, "请求参数错误！")
	} else {
		fmt.Println(article)
		ctx.String(http.StatusOK, "Success!")
	}
}
