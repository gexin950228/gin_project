package logSource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogPath  string `json:"logPath"`
	LogLevel string `json:"logLevel"`
	LogName  string `json:"logName"`
}

func LoadConfig() *LogConfig {
	logConfig := LogConfig{}
	file, err := os.Open("conf/logConfig.json")
	if err != nil {
		fmt.Printf("Open log config error, error: %v\n", err.Error())
		panic(err)
	}
	defer file.Close()
	byteData, errRead := ioutil.ReadAll(file)
	if errRead != nil {
		return nil
	}
	errUnmarshalLogConfig := json.Unmarshal(byteData, &logConfig)
	if errUnmarshalLogConfig != nil {
		return nil
	}
	return &logConfig
}
