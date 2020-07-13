package base

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"github.com/tietang/props/kvs"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	formatter *prefixed.TextFormatter
	lfh       *utils.LineNumLogrusHook
)

func init() {
	// 日志格式
	formatter = &prefixed.TextFormatter{}
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
	logrus.SetReportCaller(true)
	SetHook()
}

func SetHook() {
	lfh = utils.NewLineNumLogrusHook()
	lfh.EnableFileNameLog = true
	lfh.EnableFuncNameLog = true
	logrus.AddHook(lfh)
}

// InitLog 初始化日志
func InitLog(source kvs.ConfigSource) {
	// 设置日志输出级别
	level, err := logrus.ParseLevel(source.GetDefault("log.level", "info"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	if source.GetBoolDefault("log.enableLineLog", true) {
		lfh.EnableFuncNameLog = true
		lfh.EnableFuncNameLog = true
	} else {
		lfh.EnableFuncNameLog = false
		lfh.EnableFileNameLog = false
	}
	// 配置日志输出目录
	dir := source.GetDefault("log.dir", "./logs")
	logPath, _ := filepath.Abs(dir)
	logrus.Infof("log dir: %s", logPath)
	logFileName := source.GetDefault("log.file.name", "accesss")
	maxAge := source.GetDurationDefault("log.max.age", time.Hour*24)
	rotationTime := source.GetDurationDefault("log.rotation.time", time.Hour*1)
	os.MkdirAll(logPath, os.ModePerm)
	baseLogPath := path.Join(logPath, logFileName)

	// 设置滚动日志输出logf
	logf, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 文件切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", err)
	}

	// 设置日志文件输出的日志格式
	formatter := &logrus.TextFormatter{}
	formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		function = frame.Function
		dir, filename := path.Split(frame.File)
		f := path.Base(dir)
		return function, fmt.Sprintf("%s/%s:%d", f, filename, frame.Line)
	}

	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logf,
		logrus.InfoLevel:  logf,
		logrus.WarnLevel:  logf,
		logrus.ErrorLevel: logf,
		logrus.FatalLevel: logf,
		logrus.PanicLevel: logf,
	}, formatter)
	logrus.AddHook(hook)
}
