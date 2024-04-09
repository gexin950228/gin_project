package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var secretData map[string]interface{} = map[string]interface{}{
	"zs": gin.H{"age": 18, "addr": "张三的家"},
	"ls": gin.H{"age": 18, "addr": "李四的家"},
	"ww": gin.H{"age": "20", "addr": "王武的家"},
}

func AuthTest(ctx *gin.Context) {
	userName := ctx.Query("user_name")
	fmt.Println(userName)
	data, ok := secretData[userName]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": userName, "data": data})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{"user": userName, "data": "没有数据"})
	}
}

func WrapTest(w http.ResponseWriter, r *http.Request) {

}
