package service

import (
	"blog/internal"
	"blog/internal/dao"
	"blog/internal/dto"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type CategoryService interface {
	Select() (output dto.CategoryListOutput)
	Create(input *dto.CategoryInput) (code int)
	Update(input *dto.CategoryUpdate) bool
	Delete(id int) bool
}

type category struct {
	dao dao.CategoryDao
}

func NewCategoryService() CategoryService {
	return &category{dao: dao.NewCategoryDao()}
}

// Select 分类
func (c *category) Select() (output dto.CategoryListOutput) {
	tags, total, err := c.dao.Select(0, 100, "")
	if err != nil {
		internal.PanicCode(internal.CategoryGetFail)
	}
	if total > 0 {
		output.Total = total
		for _, item := range tags {
			output.List = append(output.List, &dto.CategoryOutput{
				ID:   item.ID,
				Name: item.Name,
			})
		}
	}
	return
}

// Create 创建分类
func (c *category) Create(input *dto.CategoryInput) (code int) {
	if c.dao.IsExist(input.Name) {
		code = internal.CategoryExist
	} else {
		_, err := c.dao.Create(&model.Category{
			Name: input.Name,
		})
		if err != nil {
			code = internal.CategoryAddFail
		}
	}
	return
}

// Update 修改分类
func (c *category) Update(input *dto.CategoryUpdate) bool {
	if c.dao.IsExist(input.Name) {
		internal.PanicCode(internal.CategoryExist)
	}
	err := c.dao.Update(&model.Category{
		Model: gorm.Model{
			ID: uint(input.ID),
		},
		Name: input.Name,
	})

	if err != nil {
		return false
	}
	return true
}

// Delete 删除分类
func (c *category) Delete(id int) bool {
	return c.dao.Delete(uint(id))
}
