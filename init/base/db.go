package base

import (
	initial "blog/init"
	"blog/internal/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"time"
)

// db 数据库实例
var db *gorm.DB

// GormDb 数据库连接池
func GormDb() *gorm.DB {
	return db
}

type settings struct {
	Driver          string
	User            string
	Password        string
	Database        string
	Host            string
	Options         map[string]string
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	LoggingEnabled  bool
}

// DbStarter 数据库starter
type DbStarter struct {
	initial.BaseStarter
}

func (s *DbStarter) Setup(ctx initial.StarterContext) {
	var err error
	conf := ctx.Props()
	setting := &settings{}
	err = kvs.Unmarshal(conf, setting, "mysql")
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(setting.Driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.User,
		setting.Password,
		setting.Host,
		setting.Database))
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(setting.MaxIdleConns)
	db.DB().SetMaxOpenConns(setting.MaxOpenConns)
	db.DB().SetConnMaxLifetime(setting.ConnMaxLifetime)

	db.LogMode(setting.LoggingEnabled)
	migrate()

	logrus.Infof("mysql.conn url:%s", setting.Host)
}

func migrate() {
	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
	}
}

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
