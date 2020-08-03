package dao

import (
	"blog/init/base"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type ArticleDao interface {
	Select(pageNum, pageSize uint, maps interface{}) (articles []model.Article, total int64, err error)
	Create(article2 *model.Article) (id uint, err error)
	Update(article2 *model.Article) (ok bool)
	Delete(id uint) bool
}

func NewArticleDao() ArticleDao {
	return &article{
		db: base.GormDb(),
	}
}

type article struct {
	db *gorm.DB
}

// Create 创建文章
func (a *article) Create(article2 *model.Article) (id uint, err error) {
	err = a.db.Create(article2).Error
	return article2.ID, err
}

// Delete 删除文章
func (a *article) Delete(id uint) bool {
	return a.db.Where("id = ?", id).Delete(&model.Article{}).RowsAffected > 0
}

// Update 更新文章
func (a *article) Update(article2 *model.Article) (ok bool) {
	return a.db.Save(article2).RowsAffected > 0
}

func (a *article) Select(pageNum, pageSize uint, maps interface{}) (articles []model.Article, total int64, err error) {
	err = a.db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Count(&total).Error
	return
}
