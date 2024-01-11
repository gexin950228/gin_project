package chapter04

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//type Article struct {
//	Id      int    `form:"-" json:"id"`
//	Title   string `form:"title" json:"title" binding:"omitempty,min=2"`
//	Author  string `form:"author" json:"author" binding:"len_valid"`
//	Content string `form:"content" json:"content" binding:"min=4,max=10"`
//	Gender  string `form:"gender" json:"gender" binding:"len_valid"`
//	Email   string `form:"email" binding:"email"`
//}

// valid for beego
type Article struct {
	Id      int    `form:"-" json:"id"`
	Title   string `form:"title" json:"title" valid:"Required"`
	Author  string `form:"author" json:"author"`
	Content string `form:"content" json:"content"`
	Gender  string `form:"gender" json:"gender"`
	Email   string `form:"email" binding:"email"`
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

func DoValidData(ctx *gin.Context) {
	var article Article
	errDataBind := ctx.ShouldBind(&article)
	if errDataBind != nil {
		fmt.Println("=============", errDataBind.Error())
	}
	valid := validation.Validation{}
	b, err1 := valid.Valid(&article)
	fmt.Println("--------", err1)
	if !b {
		for _, err2 := range valid.Errors {
			fmt.Println(err2.Message)
			ctx.String(http.StatusOK, err2.Message)
		}
	}
	ctx.String(http.StatusOK, "Success!!")
}

var LenValid validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(string)
	fmt.Println(data)
	if len(data) > 4 {
		return true
	} else {
		return false
	}
}
