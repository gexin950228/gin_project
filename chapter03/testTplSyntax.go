package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tpl1(ctx *gin.Context) {
	mapData := map[string]interface{}{
		"name": "葛新",
		"arr":  []int{1, 2, 3, 4, 5},
	}
	ctx.HTML(http.StatusOK, "chapter03/tpl1.html", mapData)
}
