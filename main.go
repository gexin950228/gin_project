package main

import (
	"fmt"
	"gin_project/chapter03"
	"gin_project/chapter04"
	_ "gin_project/dataSource"
	"gin_project/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	engine := gin.Default()
	// 创建日志文件
	file, errOpenLog := os.Create("gin_project.log")
	if errOpenLog != nil {
		fmt.Println(errOpenLog.Error())
		return
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	//engine.Use(gin.Logger(), gin.Recovery())
	//engine.Use(chapter05.MiddleWare1, chapter05.MiddleWare2())
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
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
