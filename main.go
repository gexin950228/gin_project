package main

import (
	"fmt"
	"gin_project/chapter01"
	"gin_project/chapter02"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", chapter01.Hello)

	engine.GET("/user", chapter01.User)
	engine.GET("/str", chapter02.Str)
	engine.GET("/str1", chapter02.UserInfoStruct)
	engine.GET("/arr1", chapter02.Arr)
	engine.GET("/arr2", chapter02.ArrStruct)
	engine.GET("/map", chapter02.MapController)
	engine.GET("/map1", chapter02.MapArrController)
	engine.GET("/slice", chapter02.SliceController)
	engine.GET("/param1/:id", chapter02.Param1Controller)
	// 传不传参数都可以匹配到路由
	engine.GET("/param2/*id", chapter02.Param2Controller)

	engine.GET("/param3", chapter02.GetQueryDataController)
	engine.GET("/query_array", chapter02.GetQueryArrayDataController)
	engine.GET("/query_map", chapter02.GetQueryMapDataController)
	engine.GET("/toUserAdd", chapter02.ToUserAddController)
	engine.POST("/doUserAdd", chapter02.DoUserAddController)
	engine.GET("/toUserAdd1", chapter02.ToUserAddAJAXController)
	engine.POST("/doUserAdd1", chapter02.DoUserAddAJAXController)

	engine.GET("/toSingleUpload", chapter02.ToSingleFileUploadController)
	engine.POST("/doSingleUpload", chapter02.DoSingleFileUploadController)
	engine.GET("/toSingleUploadAJAX", chapter02.ToSingleFileUploadAJAXController)
	engine.POST("/doSingleUploadAJAX", chapter02.DoSingleFileUploadAJAXController)
	engine.GET("/toMultiFleUpload", chapter02.ToMultiFileUploadController)
	engine.POST("/doMultiUpload", chapter02.DoMultiFileUploadController)
	engine.GET("/toMultiFleUploadAJAX", chapter02.ToMultiFileUploadAJAXController)
	engine.POST("/doMultiUploadAJAX", chapter02.DoMultiFileUploadAJAXController)

	engine.LoadHTMLGlob("template/**/*")
	engine.Static("/static", "static")
	err := engine.Run(":8080")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
