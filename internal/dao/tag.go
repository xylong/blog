package dao

import (
	"blog/init/base"
	"github.com/jinzhu/gorm"
)

type TagDao interface {
	Create()
	Delete()
	Update()
	Select()
}

func NewTagDao() TagDao {
	return &tag{
		db: base.GormDb(),
	}
}

type tag struct {
	db *gorm.DB
}

func (t *tag) Create() {

}

func (t *tag) Delete() {

}

func (t *tag) Update() {

}

func (t *tag) Select() {

}
