package dataSource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogMode  bool   `json:"logMode"`
}

func LoadMysqlConfig() *MysqlConfig {
	var mysqlConfig MysqlConfig
	file, errOpenFile := os.Open("../conf/mysql.conf")
	if errOpenFile != nil {
		return nil
	}
	defer file.Close()
	byteData, errReadFile := ioutil.ReadAll(file)
	if errReadFile != nil {
		return nil
	}
	errUnmarshal := json.Unmarshal(byteData, &mysqlConfig)
	if errUnmarshal != nil {
		panic(errUnmarshal.Error())
	}
	return &mysqlConfig
}
