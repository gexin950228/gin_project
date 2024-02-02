package chapter06

import (
	"gin_project/dataSource"
	"gin_project/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GormTest(ctx *gin.Context) {
	user := model.User{Name: "徐东泽", Pic: "徐东泽.png", Age: 26, Address: "哈尔滨市"}
	dataSource.Db.AutoMigrate(&model.User{})
	dataSource.Db.Create(&user)
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(dataSource.Db)
	ctx.String(http.StatusOK, "插入成功\n")
}
