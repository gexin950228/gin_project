package main

import (
	"gin_project/chapter01"
	"gin_project/chapter02"
	"gin_project/chapter03"
	"gin_project/chapter04"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	chapt01 := router.Group("/chapter01")
	chapt02 := router.Group("/chapter02")
	chapt03 := router.Group("/chapter03")
	chapt04 := router.Group("/chapter04")

	chapter01.Router(chapt01)
	chapter02.Router(chapt02)
	chapter03.Router(chapt03)
	chapter04.Router(chapt04)
}
