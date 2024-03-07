package chapter09

import (
	"fmt"
	"gin_project/logSource"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SessionTest(ctx *gin.Context) {
	// 初始化session对象
	session := sessions.Default(ctx)
	session.Set("age", "29")
	logSource.Log.Info("set session")

	name := session.Get("age")
	fmt.Printf("age: %s。\n", name)
	ctx.String(http.StatusOK, "session test!")
}
