package common

import (
	"bufio"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	conf "go-book-center/app/config"
	"os"
	"path"
	"strings"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	config := conf.Conf.Log

	if !config.IsDebug {
		logPath := config.DirName
		logFileName := config.FileName

		fileName := path.Join(logPath, logFileName)

		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)

		if err != nil {
			Logger.Errorf("Logger: open file error:", err)
		}
		defer func(src *os.File) {
			err := src.Close()
			if err != nil {
				Logger.Errorf("Logger: close file error:", err)
			}
		}(src)

		//设置输出 带缓冲批量写入
		Logger.SetOutput(bufio.NewWriter(src))
		Logger.SetLevel(logrus.InfoLevel)
		Logger.SetFormatter(&logrus.JSONFormatter{})
		writeFile(fileName, time.Duration(config.MaxAge))
	} else {
		Logger.SetLevel(logrus.DebugLevel)
	}

}

func writeFile(filename string, maxAge time.Duration) {
	logWriter, _ := rotatelogs.New(
		strings.Replace(filename, ".log", "_", -1)+".%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(maxAge*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: TIME_FORMAT,
	}))
}
