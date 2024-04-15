package model

type User struct {
	Id      int    `json:"id" form:"id" gorm:"id"`
	Name    string `json:"name" form:"name" gorm:"name"`
	Age     int    `form:"age" json:"age" gorm:"age"`
	Address string `form:"address" json:"address" gorm:"address"`
	Pic     string `form:"pic" json:"pic" gorm:"pic"`
}

type Books struct {
	Id     int    `json:"id" form:"id" gorm:"id"`
	Name   string `form:"name" json:"name" gorm:"name"`
	Author string `form:"author" json:"author" gorm:"author"`
}
