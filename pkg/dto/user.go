package dto

// RegisterInput 注册
type RegisterInput struct {
	LoginInput
	Name  string `form:"name" validate:"required,alphanum"`
	Phone string `form:"phone"`
}

// LoginInput 登录
type LoginInput struct {
	Input
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=15"`
}
