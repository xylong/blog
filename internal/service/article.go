package service

import (
	"blog/internal"
	"blog/internal/dao"
	"blog/internal/dto"
	"blog/internal/model"
	"github.com/jinzhu/gorm"
)

type ArticleService interface {
	Select(input *dto.ArticleListInput) (output dto.ArticleListOutput)
	Find(id uint) (article *dto.ArticleOutput)
	Create(input *dto.ArticleInput) (code int)
	Update(input *dto.ArticleUpdateInput) bool
	Delete(id uint) bool
}

type article struct {
	dao dao.ArticleDao
}

func NewArticleService() ArticleService {
	return &article{dao: dao.NewArticleDao()}
}

// Select 文章列表
func (a *article) Select(input *dto.ArticleListInput) (output dto.ArticleListOutput) {
	articles, total, err := a.dao.Select(input.PageNo, input.PageSize, input)
	if err != nil {
		internal.PanicCode(internal.ArticleGetFail)
	}
	if total > 0 {
		output.Total = total
		for _, item := range articles {
			output.List = append(output.List, &dto.ArticleItemOutput{
				ID:        item.ID,
				Title:     item.Title,
				CreatedAt: item.CreatedAt,
				Visits:    item.Visits,
			})
		}
	} else {
		output.List = []*dto.ArticleItemOutput{}
	}
	return
}

// Find 文章详情
func (a *article) Find(id uint) (article *dto.ArticleOutput) {
	result, err := a.dao.Find(id)
	if err != nil {
		internal.PanicCode(internal.ArticleGetFail)
	}
	article = &dto.ArticleOutput{
		ArticleItemOutput: dto.ArticleItemOutput{
			ID:        result.ID,
			Title:     result.Title,
			CreatedAt: result.CreatedAt,
			Visits:    result.Visits,
		},
		CategoryID:   result.CategoryId,
		CategoryName: result.Category.Name,
		Content:      result.Content,
	}

	return
}

// Create 写文章
func (a *article) Create(input *dto.ArticleInput) (code int) {
	_, err := a.dao.Create(&model.Article{
		CategoryId: uint(input.CategoryId),
		Title:      input.Title,
		Content:    input.Content,
		Visits:     1,
	})
	if err != nil {
		code = internal.ArticleAddFail
	}
	return
}

// Update 修改文章
func (a *article) Update(input *dto.ArticleUpdateInput) bool {
	return a.dao.Update(&model.Article{
		Model: gorm.Model{
			ID: uint(input.ID),
		},
		CategoryId: uint(input.CategoryId),
		Title:      input.Title,
		Content:    input.Content,
	})
}

// Delete 删除文章
func (a *article) Delete(id uint) bool {
	return a.dao.Delete(id)
}
