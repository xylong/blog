package dao

import (
	"blog/init/base"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type UserDao interface {
	Find(*model.User) (*model.User, error)
	Create(*model.User) (uint, error)
	IsExist(email, phone string) bool
}

func NewUserDao() UserDao {
	return &user{
		db: base.GormDb(),
	}
}

type user struct {
	db *gorm.DB
}

// Find 查找用户
func (dao *user) Find(user *model.User) (*model.User, error) {
	err := dao.db.Where(user).First(user).Error
	return user, err
}

// Create 创建用户
func (dao *user) Create(user *model.User) (uint, error) {
	err := dao.db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// IsExist 判断用户是否存在
func (dao *user) IsExist(email, phone string) bool {
	var user model.User
	if email != "" && phone != "" {
		dao.db.Where("email = ?", email).Or("phone = ?", phone).First(&user)
	} else if email != "" && phone == "" {
		dao.db.Where("email = ?", email).First(&user)
	} else if email == "" && phone != "" {
		dao.db.Where("phone = ?", phone).First(&user)
	}
	return user.ID > 0
}
