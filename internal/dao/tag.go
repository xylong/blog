package dao

import (
	"blog/init/base"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type TagDao interface {
	Create(*model.Tag) (uint, error)
	Delete(uint) bool
	Update(*model.Tag) error
	Select(uint, uint, interface{}) ([]model.Tag, int64, error)
	// IsExist 判断标签是否存在
	IsExist(string) bool
}

func NewTagDao() TagDao {
	return &tag{
		db: base.GormDb(),
	}
}

type tag struct {
	db *gorm.DB
}

// Create 创建标签
func (t *tag) Create(tag *model.Tag) (uint, error) {
	err := t.db.Create(tag).Error
	if err != nil {
		return 0, err
	}
	return tag.ID, nil
}

// Delete 删除标签
func (t *tag) Delete(id uint) bool {
	return t.db.Where("id = ?", id).Delete(&model.Tag{}).RowsAffected > 0
}

// Update 更新标签
func (t *tag) Update(tag *model.Tag) (err error) {
	return t.db.Save(tag).Error
}

// Select 查询标签
func (t *tag) Select(pageNum, pageSize uint, maps interface{}) (tags []model.Tag, total int64, err error) {
	err = t.db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Count(&total).Error
	return
}

// IsExist 标签是否存在
func (t *tag) IsExist(name string) bool {
	var tag1 model.Tag
	t.db.Where("name = ?", name).First(&tag1)
	return tag1.ID > 0
}
