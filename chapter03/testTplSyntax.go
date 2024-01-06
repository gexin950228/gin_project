package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id    int
	Title string
	Desc  string
}

func Tpl1(ctx *gin.Context) {
	article := Article{
		Id:    1,
		Title: "西游记",
		Desc:  "西天取经",
	}
	mapData := map[string]interface{}{
		"name":    "葛新",
		"arr":     []int{1, 2, 3, 4, 5},
		"article": article,
	}
	ctx.HTML(http.StatusOK, "chapter03/tpl1.html", mapData)
}
