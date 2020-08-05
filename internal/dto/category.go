package dto

type CategoryInput struct {
	Input
	Name string `json:"name" form:"name" comment:"分类名" example:"编程" validate:"required,alphanumunicode"`
}

type CategoryOutput struct {
	ID   uint   `json:"id"`   // id
	Name string `json:"name"` // 分类名
}

type CategoryListOutput struct {
	List  []*CategoryOutput `json:"list" comment:"分类列表"`
	Total int64             `json:"total" comment:"分类数量"`
}

type CategoryUpdate struct {
	CategoryInput
	ID int `json:"id" form:"id" comment:"分类id" example:"1" validate:"required,gt=0"` // id
}
