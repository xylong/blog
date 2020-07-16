package main

import (
	initial "blog/init"
	"blog/init/base"
	"blog/internal/router"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"os"
)

func init() {
	initial.Register(&base.PropsStarter{})
	initial.Register(&base.DbStarter{})
}

func main() {
	// 设置配置文件路径
	dir, _ := os.Getwd()
	viper.SetConfigName("mysql")                             // 配置文件名
	viper.SetConfigType("toml")                              // 配置扩展名
	viper.AddConfigPath(fmt.Sprintf("%s/%s", dir, "config")) // 查找配置文件所在路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic("未找到配置文件")
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic("配置文件内容错误")
		}
	}
	file := kvs.GetCurrentFilePath("config/config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := initial.NewBoot(conf)
	app.Start()

	engine := router.InitRouter()
	engine.Run()
}
