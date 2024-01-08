package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int    `form:"id" json:"id" uri:"id"`
	Name string `form:"name" json:"name" uri:"name"`
	Addr string `form:"addr" json:"addr" uri:"addr"`
	Age  string `form:"age" json:"age" uri:"age"`
}

func ShouldBind1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind1.html", nil)
}

func DoShouldBind1(ctx *gin.Context) {
	var user UserInfo
	errBind := ctx.ShouldBind(&user)
	if errBind != nil {
		fmt.Println(errBind.Error())
		return
	} else {
		fmt.Println(user)
	}
	ctx.String(http.StatusOK, "Success!")
}

func ShouldBind2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind2.html", nil)
}

func DoShouldBind2(ctx *gin.Context) {
	var user UserInfo
	errJsonBind := ctx.ShouldBindJSON(&user)
	var data map[string]interface{}
	if errJsonBind != nil {
		data = map[string]interface{}{
			"code": 500,
			"msg":  "JSON数据绑定出错！",
		}
	} else {
		data = map[string]interface{}{
			"code": 200,
			"msg":  "JSON数据提交成功！",
		}
	}
	fmt.Println(user)
	ctx.JSON(http.StatusOK, data)
}

func BindUri(ctx *gin.Context) {
	var user UserInfo
	err := ctx.ShouldBindUri(&user)
	fmt.Println(user)
	if err != nil {
		return
	}
	fmt.Println(user)
	ctx.String(http.StatusOK, "Success")
}
