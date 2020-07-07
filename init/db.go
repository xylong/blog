package init

import (
	"blog/pkg/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// db 数据库实例
var db *gorm.DB

// GormDb 数据库连接池
func GormDb() *gorm.DB {
	return db
}

func InitDb() {
	var err error
	db, err = gorm.Open(viper.GetString("database.driver"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.name")))

	if err != nil {
		panic("connect mysql failed, " + err.Error())
	}
}

// DbStarter 数据库starter
type DbStarter struct {
	util.BaseStarter
}
