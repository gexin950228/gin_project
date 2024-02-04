package chapter07

import (
	"gin_project/dataSource"
	"gin_project/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GormTest(ctx *gin.Context) {
	user := model.User{Name: "葛新", Pic: "葛新.png", Age: 29, Address: "湘潭市"}
	dataSource.Db.AutoMigrate(&model.User{})
	dataSource.Db.Create(&user)
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(dataSource.Db)
	ctx.String(http.StatusOK, "插入成功\n")
}
