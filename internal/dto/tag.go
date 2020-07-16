package dto

type TagInput struct {
	Input
	Name string `json:"name" form:"name" comment:"标签名" example:"go" validate:"required,alphanumunicode"`
}
