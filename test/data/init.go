package data

import (
	initial "blog/init"
	"blog/init/base"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func init() {
	initial.Register(&base.PropsStarter{})
	initial.Register(&base.DbStarter{})

	file := kvs.GetCurrentFilePath("../../config/config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := initial.NewBoot(conf)
	app.Start()
}
