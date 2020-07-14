package dao

import (
	"blog/init/base"
	"blog/pkg/model"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Find(*model.User) (*model.User, error)
	Create(*model.User) (uint, error)
	IsExist(email, phone string) bool
}

func NewDao() Dao {
	return &UserDao{
		db: base.GormDb(),
	}
}

type UserDao struct {
	db *gorm.DB
}

// Find 查找用户
func (dao *UserDao) Find(user *model.User) (*model.User, error) {
	err := dao.db.Where(user).First(user).Error
	return user, err
}

// Create 创建用户
func (dao *UserDao) Create(user *model.User) (uint, error) {
	err := dao.db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// IsExist 判断用户是否存在
func (dao *UserDao) IsExist(email, phone string) bool {
	var user model.User
	dao.db.Where("email = ?", email).Or("phone = ?", phone).First(&user)
	return user.ID > 0
}
