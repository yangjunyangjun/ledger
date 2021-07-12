package sdk

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Log *log.Logger

func newLog(logFileName string) {
	logPath, _ := os.Getwd()
	fileName := path.Join(logPath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		src, err = os.Create(fileName)
		if err != nil {
			panic(fmt.Sprintf("log path error,path:%s", fileName))
		}
	}
	Log = log.New()
	Log.Out = src
	//设置日志等级
	Log.SetLevel(log.DebugLevel)
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("logWriter error:%s", err.Error()))
	}
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	//设置日志格式
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	log.AddHook(lfHook)

}
