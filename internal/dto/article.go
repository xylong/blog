package dto

import "time"

type ArticleInput struct {
	Input
	CategoryId int    `json:"category_id" form:"category_id" comment:"分类id" example:"1" validate:"required,numeric"`
	Title      string `json:"title" form:"title" comment:"标题" example:"世界辣么大" validate:"required,alphanumunicode"`
	Content    string `json:"content" form:"content" comment:"内容" example:"我想去看看" validate:"required,alphanumunicode"`
}

type ArticleListInput struct {
	Input
	CategoryId int    `json:"category_id" form:"category_id" comment:"分类id" example:"1" validate:"omitempty,numeric,gt=0"` // 分类id
	Title      string `json:"title" form:"title" comment:"标题" example:"世界辣么大" validate:"omitempty,alphanum"`               // 搜索关键字
	PageSize   int    `json:"page_size" form:"page_size" comment:"条数" validate:"required,min=1,max=100"`                   // 每页条数
	PageNo     int    `json:"page_no" form:"page_no" comment:"页码" validate:"required,gt=0"`                                // 页码
}

type ArticleItemOutput struct {
	ID        uint      `json:"id"`         // id
	Title     string    `json:"title"`      // 标题
	Visits    uint      `json:"visits"`     // 访问量
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

type ArticleListOutput struct {
	List  []*ArticleItemOutput `json:"list" comment:"标签列表"`  // 列表
	Total int64                `json:"total" comment:"标签数量"` // 数量
}

type ArticleOutput struct {
	ArticleItemOutput
	CategoryID   uint   `json:"category_id"`   // 分类id
	CategoryName string `json:"category_name"` // 分类name
	Content      string `json:"content"`       // 内容
}

type ArticleUpdateInput struct {
	Input
	ID         int    `json:"id" form:"id" comment:"文章id" example:"1" validate:"required,gt=0"` // id
	CategoryId int    `json:"category_id" form:"category_id" comment:"分类id" example:"1" validate:"omitempty,numeric"`
	Title      string `json:"title" form:"title" comment:"标题" example:"世界辣么大" validate:"omitempty,alphanumunicode"`
	Content    string `json:"content" form:"content" comment:"内容" example:"我想去看看" validate:"omitempty,alphanumunicode"`
}
