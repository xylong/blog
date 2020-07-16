package dto

type CategoryInput struct {
	Input
	Name string `json:"name" form:"name" comment:"分类名" example:"编程" validate:"required,alphanumunicode"`
}
