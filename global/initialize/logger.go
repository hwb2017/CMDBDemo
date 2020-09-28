package initialize

import (
	"fmt"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/utils"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const LOGNAME = "cmdb"

func InitLogger() {
	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Level = setupLogLevel()
	logger.Out = setupLogPath()
	global.Logger = logger
}

func setupLogLevel() logrus.Level {
	switch global.LogConfiguration.Level {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func setupLogPath() io.Writer {
	logPath := global.LogConfiguration.Path
	if logPath == "" {
		return os.Stdout
	}
	if ok, _ := utils.PathExists(logPath); !ok {
		fmt.Println("create log directory")
		os.Mkdir(logPath, 0666)
	}
	logFullPath := logPath+string(os.PathSeparator)+LOGNAME+".log"
	fmt.Println(logFullPath)
	fileOutput, err := os.OpenFile(logFullPath, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to log to file, using default stdout")
		return os.Stdout
	}
	return fileOutput
}