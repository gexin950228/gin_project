package main

import (
	"fmt"
	"gin_project/logSource"
	"github.com/gin-contrib/sessions"

	//"fmt"
	"gin_project/chapter03"
	"gin_project/chapter04"
	_ "gin_project/dataSource"
	_ "gin_project/logSource"
	"gin_project/router"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	//store := cookie.NewStore([]byte("Hello"))
	//engine.Use(sessions.Sessions("gin_session", store))

	// redis为存储引擎的session初始化
	store, errInitStore := redis.NewStore(10, "tcp", "localhost:6379", "Gexin..950228", []byte("secret"))
	if errInitStore != nil {
		fmt.Println(errInitStore.Error())
		logSource.Log.Error("初始化session失败")
		fmt.Println(errInitStore.Error())
		return
	} else {
		fmt.Println("初始化session成功")
		logSource.Log.Info("初始化session成功")
		engine.Use(sessions.Sessions("gin_session", store))
	}

	//-----------------------------------------------------------------------
	// 创建日志文件
	// file, errOpenLog := os.OpenFile("gin_project.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if errOpenLog != nil {
	//	fmt.Println(errOpenLog.Error())
	//	return
	//}
	// gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	//engine.Use(gin.Logger(), gin.Recovery())
	//engine.Use(middleware.MiddleWare1, middleware.MiddleWare2())
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("len_valid", chapter04.LenValid)
	}
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
	router.Router(engine)
	s := &http.Server{
		Addr:         ":8000",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
