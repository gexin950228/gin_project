package dataSource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {
	fmt.Println("================================================")
	dataSource := LoadMysqlConfig()
	fmt.Println(dataSource)
	fmt.Println("==================================================")
	Db, err = gorm.Open("mysql", "root:Gexin..950228@tcp(localhost:3306)/gin_project?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("数据库连接成功")
	Db.LogMode(true)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(50)
	//Db.AutoMigrate(&model.User{})
}
