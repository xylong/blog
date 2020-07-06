package base

import (
	"blog/pkg/util"
	"github.com/jinzhu/gorm"
)

// db 数据库实例
type db *gorm.DB

// GormDb 数据库连接池
func GormDb() *gorm.DB {
	return db
}

// DbStarter 数据库starter
type DbStarter struct {
	util.BaseStarter
}
