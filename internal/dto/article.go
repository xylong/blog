package dto

type ArticleInput struct {
	Input
	CategoryId uint   `json:"category_id" form:"category_id" comment:"分类id" example:"1" validate:"required,numeric"`
	TagId      uint   `json:"tag_id" form:"tag_id" comment:"标签id" example:"1" validate:"numeric"`
	Title      string `json:"title" form:"title" comment:"标题" example:"世界辣么大" validate:"required,alphanumunicode"`
	Content    string `json:"content" form:"content" comment:"内容" example:"我想去看看" validate:"required,alphanumunicode"`
	Visits     uint   `json:"visits" form:"visits" comment:"浏览次数" example:"100" validate:"required,numeric"`
}
