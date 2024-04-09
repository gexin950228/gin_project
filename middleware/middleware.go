package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare1(ctx *gin.Context) {
	startTime := time.Now()
	fmt.Println("这是自定义中间件1...1")
	ctx.Next()
	fmt.Println("这是自定义中间件1...2")
	endTime := time.Since(startTime)
	fmt.Printf("耗时：%v\ns", endTime)
}

func MiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		time.Sleep(200)
		fmt.Println("这是自定义中间件2")
	}
}

// Cors 解决跨域
func Cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Max-Age", "86400")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusFound)
	}
	ctx.Next()
}
