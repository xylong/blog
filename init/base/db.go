package base

import (
	"github.com/jinzhu/gorm"
)

// db 数据库实例
var db *gorm.DB

// GormDb 数据库连接池
func GormDb() *gorm.DB {
	return db
}

// DbStarter 数据库starter
//type DbStarter struct {
//	init.BaseStarter
//}
//
//func (s *DbStarter) Setup() {
//	var err error
//	db, err = gorm.Open(viper.GetString("database.driver"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		viper.GetString("database.user"),
//		viper.GetString("database.password"),
//		viper.GetString("database.host"),
//		viper.GetString("database.name")))
//	if err != nil {
//		panic("connect mysql failed, " + err.Error())
//	}
//
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//	db.DB().SetConnMaxLifetime(10)
//}
