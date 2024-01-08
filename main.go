package main

import (
	"gin_project/chapter01"
	"gin_project/chapter02"
	"gin_project/chapter03"
	"gin_project/chapter04"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("len_valid", chapter04.LenValid)
	}
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
	engine.GET("/json", chapter02.OutJSON)
	engine.GET("/ascii_json", chapter02.AsciiJson)
	engine.GET("/jsonp", chapter02.JsonP)
	engine.GET("/yaml", chapter02.Yaml)
	engine.GET("/xml", chapter02.Xml)
	engine.GET("/secure_json", chapter02.SecureJson)
	engine.GET("/pure_json", chapter02.PureJson)
	engine.GET("/proto", chapter02.ProtoController)

	engine.GET("/tpl1", chapter03.Tpl1)
	engine.GET("/func1", chapter03.Func1)

	engine.GET("/bind1", chapter04.ShouldBind1)
	engine.POST("/doBind1", chapter04.DoShouldBind1)
	engine.GET("/bind2", chapter04.ShouldBind2)
	engine.POST("/doBind2", chapter04.DoShouldBind2)
	engine.GET("/to_valid1", chapter04.ToValidData)
	engine.POST("/do_valid1", chapter04.DoValidData)

	engine.GET("/bind_uri/:id/:name/:addr/:age", chapter04.BindUri)

	// 注册模板函数
	engine.SetFuncMap(template.FuncMap{
		"CutStr":   chapter03.CutStr,
		"safeHtml": chapter03.SafeHTML,
	})
	engine.LoadHTMLGlob("template/**/*")
	engine.Static("/static", "static")
	//http.ListenAndServe(":8080", engine)
	//err := engine.Run(":8080")
	//now := time.Now()
	//nowFormat := now.Format("2006.01.02.15.04.5")
	//fmt.Println(nowFormat)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
