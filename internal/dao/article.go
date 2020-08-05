package dao

import (
	"blog/init/base"
	"blog/internal/dto"
	"blog/internal/model"
	"blog/pkg/util"
	"github.com/jinzhu/gorm"
)

type ArticleDao interface {
	Select(pageNum, pageSize int, maps interface{}) (articles []model.Article, total int64, err error)
	Find(id uint) (article *model.Article, err error)
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

// Select 文章查询
func (a *article) Select(pageNum, pageSize int, maps interface{}) (articles []model.Article, total int64, err error) {
	offset := util.GetPage(pageNum, pageSize)
	table := a.db.NewScope(&model.Article{}).TableName()
	query := a.db.Table(table).Select("id,title,created_at,visits")

	if v, ok := maps.(*dto.ArticleListInput); ok {
		if v.Title != "" {
			query = query.Where("title like ?", "%"+v.Title+"%")
		}
		if v.CategoryId > 0 {
			query = query.Where("category_id = ?", v.CategoryId)
		}
	}

	err = query.Offset(offset).Limit(pageSize).Order("id desc").Find(&articles).Count(&total).Error
	return
}

// Find 单个查找
func (a *article) Find(id uint) (article *model.Article, err error) {
	article = &model.Article{}
	err = a.db.Preload("Category").First(article, id).Error
	return
}
