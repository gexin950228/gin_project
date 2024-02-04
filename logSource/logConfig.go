package logSource

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	logConfig := LoadConfig()

	//设置日志输出文件
	logDir := logConfig.LogPath + "/" + logConfig.LogName
	file, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	Log.Out = file
	levelMapping := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.DebugLevel,
		"warn":  logrus.WarnLevel,
		"trace": logrus.TraceLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(levelMapping[logConfig.LogLevel])
	Log.SetFormatter(&logrus.TextFormatter{})
}
