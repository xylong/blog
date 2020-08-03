package dao

import (
	"blog/init/base"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type CategoryDao interface {
	Create(category *model.Category) (id uint, err error)
	Select(pageNum, pageSize uint, maps interface{}) (categories []model.Category, total int64, err error)
	Update(category *model.Category) (err error)
	Delete(uint) bool
	IsExist(string) bool
}

func NewCategoryDao() CategoryDao {
	return &category{
		db: base.GormDb(),
	}
}

type category struct {
	db *gorm.DB
}

// Create 创建分类
func (c *category) Create(category *model.Category) (id uint, err error) {
	err = c.db.Create(category).Error
	return category.ID, err
}

// Delete 删除分类
func (c *category) Delete(id uint) bool {
	return c.db.Where("id = ?", id).Delete(&model.Category{}).RowsAffected > 0
}

// Update 更改分类
func (c *category) Update(category *model.Category) (err error) {
	return c.db.Save(category).Error
}

// Select 查询分类
func (c *category) Select(pageNum, pageSize uint, maps interface{}) (categories []model.Category, total int64, err error) {
	err = c.db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories).Count(&total).Error
	return
}

// IsExist 标签是否存在
func (c *category) IsExist(name string) bool {
	var category1 model.Category
	c.db.Where("name = ?", name).First(&category1)
	return category1.ID > 0
}
