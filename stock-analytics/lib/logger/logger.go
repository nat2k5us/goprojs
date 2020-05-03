package logger

import (
	"strings"

	"github.com/marvin5064/stock-analytics/lib/utils"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewDevelopment()
)

func Sync() {
	logger.Sync()
}
func Info(input ...interface{}) {
	details := utils.ArrayToStrings(input)
	if len(details) == 0 {
		return
	}
	logger.Info(strings.Join(details, " "))
}

func Warn(input ...interface{}) {
	details := utils.ArrayToStrings(input)
	if len(details) == 0 {
		return
	}
	logger.Warn(strings.Join(details, " "))
}

func Error(input ...interface{}) {
	details := utils.ArrayToStrings(input)
	if len(details) == 0 {
		return
	}
	logger.Error(strings.Join(details, " "))
}

func Fatal(input ...interface{}) {
	details := utils.ArrayToStrings(input)
	if len(details) == 0 {
		return
	}
	logger.Fatal(strings.Join(details, " "))
}
