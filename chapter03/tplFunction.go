package chapter03

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Func1(ctx *gin.Context) {
	mapData := map[string]interface{}{
		"age":        28,
		"name":       "葛新",
		"slice_data": []int{1, 2, 3, 4, 5},
	}
	ctx.HTML(http.StatusOK, "chapter03/func.html", mapData)
}

func CutStr(str string, num int) string {
	if len(str) <= num+3 {
		return str
	} else {
		return str[:num] + "..."
	}
}

func SafeHTML(str string) template.HTML {
	return template.HTML(str)
}
