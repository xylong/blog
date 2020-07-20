package dao

import (
	"blog/init/base"
	"github.com/jinzhu/gorm"
)

type CategoryDao interface {
	Create()
	Select()
	Update()
	Delete()
}

func NewCategoryDao() CategoryDao {
	return &article{
		db: base.GormDb(),
	}
}

type category struct {
	db *gorm.DB
}

func (c *category) Create() {

}

func (c *category) Delete() {

}

func (c *category) Update() {

}

func (c *category) Select() {

}
