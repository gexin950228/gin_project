package chapter11

import (
	"fmt"
	"gin_project/dataSource"
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
	var books []model.Books
	err := dataSource.Db.Find(&books).Error
	if err != nil {
		fmt.Println("查询数据出错")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func GetBookDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	var book Book
	dataSource.Db.Debug().First(&book, "id=?", id)
	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func PostTest(ctx *gin.Context) {
	name := ctx.PostForm("name")
	fmt.Println(name)
	data := map[string]interface{}{
		"code": 1,
		"msg":  "提交成功",
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func FileUpload(ctx *gin.Context) {
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(name, password)
	uploadFile, _ := ctx.FormFile("upload_file")
	fileName := uploadFile.Filename
	filePath := "upload/" + fileName
	err := ctx.SaveUploadedFile(uploadFile, filePath)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "Success",
	})
}

func FilesUpload(ctx *gin.Context) {
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(name, password)
	uploadFiles, _ := ctx.MultipartForm()
	files := uploadFiles.File["upload_files"]
	for _, file := range files {
		filePath := "upload/" + file.Filename
		err := ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "Success",
	})
}
