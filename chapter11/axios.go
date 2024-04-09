package chapter11

import (
	"gin_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiAxios(ctx *gin.Context) {
	// 结构体
	user := model.User{
		Id:      1,
		Name:    "葛新",
		Address: "北京市昌平区",
		Age:     29,
		Pic:     "gexin.png",
	}
	array := [4]int{1, 2, 3, 4}
	// map
	// 结构体数组
	// slice
	// 结构体map
	arrayStruct := []model.User{
		{Id: 1, Name: "葛新", Address: "北京市昌平区", Pic: "gexin.png", Age: 29},
		{Id: 2, Name: "王岩", Address: "北京市朝阳区", Pic: "wangyan.png", Age: 28},
	}
	mapStruct := map[string]model.User{
		"user": model.User{Id: 3, Name: "满志伟", Address: "北京市", Pic: "mzw.png", Age: 26},
	}
	data := map[string]interface{}{
		"code":        "200",
		"msg":         "成功",
		"user":        user,
		"arrays":      array,
		"arrayStruct": arrayStruct,
		"mapStruct":   mapStruct,
	}

	ctx.JSON(http.StatusOK, data)
}

type Book struct {
	Id     int    `json:"id" form:"id"`
	Name   string `form:"name" json:"name"`
	Author string `form:"author" json:"author"`
}

func GetBooks(ctx *gin.Context) {
	books := []Book{
		{Id: 1, Name: "水浒传", Author: "施耐庵"},
		{Id: 2, Name: "西游记", Author: "吴承恩"},
		{Id: 3, Name: "红楼梦", Author: "曹雪芹"},
		{Id: 4, Name: "三国", Author: "罗贯中"},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
