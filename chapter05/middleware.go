package chapter05

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
