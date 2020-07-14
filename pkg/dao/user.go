package dao

import (
	"blog/init/base"
	"blog/pkg/model"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Find(*model.User) error
	Create(*model.User) (uint, error)
}

func NewDao() Dao {
	return &UserDao{db: base.GormDb()}
}

type UserDao struct {
	db *gorm.DB
}

// Find 查找用户
func (dao *UserDao) Find(user *model.User) error {
	return dao.db.First(user).Error
}

// Create 创建用户
func (dao *UserDao) Create(user *model.User) (uint, error) {
	err := dao.db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
