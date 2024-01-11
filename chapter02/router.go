package chapter02

import "github.com/gin-gonic/gin"

func Router(chapt02 *gin.RouterGroup) {
	chapt02.GET("/str", Str)
	chapt02.GET("/str1", UserInfoStruct)
	chapt02.GET("/arr1", Arr)
	chapt02.GET("/arr2", ArrStruct)
	chapt02.GET("/map", MapController)
	chapt02.GET("/map1", MapArrController)
	chapt02.GET("/slice", SliceController)
	chapt02.GET("/param1/:id", Param1Controller)
	// 传不传参数都可以匹配到路由
	chapt02.GET("/param2/*id", Param2Controller)

	chapt02.GET("/param3", GetQueryDataController)
	chapt02.GET("/query_array", GetQueryArrayDataController)
	chapt02.GET("/query_map", GetQueryMapDataController)
	chapt02.GET("/toUserAdd", ToUserAddController)
	chapt02.POST("/doUserAdd", DoUserAddController)
	chapt02.GET("/toUserAdd1", ToUserAddAJAXController)
	chapt02.POST("/doUserAdd1", DoUserAddAJAXController)

	chapt02.GET("/toSingleUpload", ToSingleFileUploadController)
	chapt02.POST("/doSingleUpload", DoSingleFileUploadController)
	chapt02.GET("/toSingleUploadAJAX", ToSingleFileUploadAJAXController)
	chapt02.POST("/doSingleUploadAJAX", DoSingleFileUploadAJAXController)
	chapt02.GET("/toMultiFleUpload", ToMultiFileUploadController)
	chapt02.POST("/doMultiUpload", DoMultiFileUploadController)
	chapt02.GET("/toMultiFleUploadAJAX", ToMultiFileUploadAJAXController)
	chapt02.POST("/doMultiUploadAJAX", DoMultiFileUploadAJAXController)
	chapt02.GET("/json", OutJSON)
	chapt02.GET("/ascii_json", AsciiJson)
	chapt02.GET("/jsonp", JsonP)
	chapt02.GET("/yaml", Yaml)
	chapt02.GET("/xml", Xml)
	chapt02.GET("/secure_json", SecureJson)
	chapt02.GET("/pure_json", PureJson)
	chapt02.GET("/proto", ProtoController)
}
