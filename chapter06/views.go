package chapter06

import (
	"gin_project/dataSource"
	"gin_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GormTest(ctx *gin.Context) {
	user := model.User{Id: 31, Name: "徐东泽", Pic: "徐东泽.png", Age: 26, Address: "哈尔滨市"}
	dataSource.Db.Create(&user)
	dataSource.Db.Close()
	ctx.String(http.StatusOK, "插入成功\n")
}
