package dto

type TagInput struct {
	Input
	Name string `json:"name" form:"name" comment:"标签名" example:"mysql" validate:"required,alphanumunicode"` // 标签名
}

type TagUpdate struct {
	TagInput
	ID int `json:"id" form:"name" comment:"标签id" example:"1" validate:"required,gt=0"` // id
}

type TagOutput struct {
	ID   uint   `json:"id"`   // id
	Name string `json:"name"` // 标签名
}

type TagListOutput struct {
	List  []*TagOutput `json:"list" comment:"标签列表"`
	Total int64        `json:"total" comment:"标签数量"`
}
