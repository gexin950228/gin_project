package chapter02

import (
	user "gin_project/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OutJSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func AsciiJson(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func SecureJson(ctx *gin.Context) {
	ctx.SecureJSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func Xml(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func Yaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
func JsonP(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func PureJson(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	//	---------------------------------------------------------
}

func ProtoController(ctx *gin.Context) {
	user1 := &user.User{Name: "John", Age: 29}
	ctx.ProtoBuf(http.StatusOK, user1)
}
