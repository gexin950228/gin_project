package dataSource

import (
	"fmt"
	"gin_project/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {
	dataSource := LoadMysqlConfig()
	host := dataSource.Host
	port := dataSource.Port
	user := dataSource.User
	password := dataSource.Password
	database := dataSource.DataBase
	logMode := dataSource.LogMode
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user, password, host, port, database,
	)
	fmt.Printf("dst: %v\n", dst)
	Db, err = gorm.Open("mysql", dst)
	Db.LogMode(logMode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("数据库连接成功")
	Db.LogMode(true)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(50)
	Db.AutoMigrate(&model.Books{})
}
