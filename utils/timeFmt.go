package utils

import (
	"fmt"
	"strings"
	"time"
)

func FileNameFormat(str string) string {
	now := time.Now()
	timeStr := fmt.Sprintf(now.Format("2006-01-01-15-04-05"))
	parts := strings.Split(str, ".")
	var fileName = ""
	for i := 0; i < (len(parts) - 1); i++ {
		fileName = fileName + "." + parts[i]
	}
	fileName = strings.TrimLeft(fileName+"-"+timeStr+"."+parts[len(parts)-1], ".")
	return fileName
}
