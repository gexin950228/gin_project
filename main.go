package main

import (
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
	Router(engine)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
