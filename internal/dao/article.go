package dao

import (
	"blog/init/base"
	"github.com/jinzhu/gorm"
)

type ArticleDao interface {
	Create()
	Select()
	Update()
	Delete()
}

func NewArticleDao() ArticleDao {
	return &article{
		db: base.GormDb(),
	}
}

type article struct {
	db *gorm.DB
}

func (a *article) Create() {

}

func (a *article) Delete() {

}

func (a *article) Update() {

}

func (a *article) Select() {

}
