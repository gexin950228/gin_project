package router

import (
	"gin_project/chapter01"
	"gin_project/chapter02"
	"gin_project/chapter03"
	"gin_project/chapter04"
	"gin_project/chapter05"
	"gin_project/chapter06"
	"gin_project/chapter07"
	"gin_project/chapter09"
	"gin_project/chapter11"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	chapt01 := router.Group("/chapter01")
	chapt02 := router.Group("/chapter02")
	chapt02.Use(chapter05.MiddleWare1, chapter05.MiddleWare2())
	chapt03 := router.Group("/chapter03")
	chapt04 := router.Group("/chapter04")
	chapt05 := router.Group("/chapter05")
	chapt06 := router.Group("/chapter06")
	chapt07 := router.Group("/chapter07")
	chapt09 := router.Group("/chapter09")
	chapt11 := router.Group("/chapter11")
	chapter01.Router(chapt01)
	chapter02.Router(chapt02)
	chapter03.Router(chapt03)
	chapter04.Router(chapt04)
	chapter05.Router(chapt05)
	chapter06.Router(chapt06)
	chapter07.Router(chapt07)
	chapter09.Router(chapt09)
	chapter11.Router(chapt11)
}
