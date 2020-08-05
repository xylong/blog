package service

import (
	"blog/internal"
	"blog/internal/dao"
	"blog/internal/dto"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type TagService interface {
	Create(input *dto.TagInput) (code int)
	Select() (output dto.TagListOutput)
	Update(input *dto.TagUpdate) bool
	Delete(id int) bool
}

type tag struct {
	dao dao.TagDao
}

func NewTagService() TagService {
	return &tag{
		dao: dao.NewTagDao(),
	}
}

// Select 所有标签
func (t *tag) Select() (output dto.TagListOutput) {
	tags, total, err := t.dao.Select(0, 100, "")
	if err != nil {
		internal.PanicCode(internal.TagGetFail)
	}
	if total > 0 {
		output.Total = total
		for _, item := range tags {
			output.List = append(output.List, &dto.TagOutput{
				ID:   item.ID,
				Name: item.Name,
			})
		}
	} else {
		output.List = []*dto.TagOutput{}
	}
	return
}

// Create 创建标签
func (t *tag) Create(input *dto.TagInput) (code int) {
	if t.dao.IsExist(input.Name) {
		code = internal.TagExist
	} else {
		_, err := t.dao.Create(&model.Tag{
			Name: input.Name,
		})
		if err != nil {
			code = internal.TagAddFail
		}
	}
	return
}

// Update 修改标签
func (t *tag) Update(input *dto.TagUpdate) bool {
	if t.dao.IsExist(input.Name) {
		internal.PanicCode(internal.TagExist)
	}
	err := t.dao.Update(&model.Tag{
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

// Delete 删除标签
func (t *tag) Delete(id int) bool {
	return t.dao.Delete(uint(id))
}
