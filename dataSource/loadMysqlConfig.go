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
	mysqlConfig := MysqlConfig{}
	file, errOpenFile := os.Open("conf/mysql.json")
	if errOpenFile != nil {
		panic(errOpenFile)
	}
	defer file.Close()
	byteData, errReadFile := ioutil.ReadAll(file)
	if errReadFile != nil {
		panic(errReadFile)
	}
	errUnmarshal := json.Unmarshal(byteData, &mysqlConfig)
	if errUnmarshal != nil {
		panic(errUnmarshal.Error())
	}
	return &mysqlConfig
}
