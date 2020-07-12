package base

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"log"
	"os"
	"time"
)

func init() {
	// 日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000"
	formatter.ForceFormatting = true // 强制格式化
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	logrus.SetFormatter(formatter)
	// 日志级别
	level := os.Getenv("log.debug")
	if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	// 控制台高亮
	formatter.ForceColors = true
	formatter.DisableColors = false
	// 日志文件和滚动配置
	logf, err := rotatelogs.New(
		"%Y-%m-%d-%H-%M.log",
		//rotatelogs.WithLinkName("/path/to/access_log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),  // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour), // 文件切割时间间隔
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	// 文件输出格式
	fileFormatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}
	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logf,
		logrus.InfoLevel:  logf,
		logrus.WarnLevel:  logf,
		logrus.ErrorLevel: logf,
		logrus.FatalLevel: logf,
		logrus.PanicLevel: logf,
	}, fileFormatter)
	logrus.AddHook(hook)
}
