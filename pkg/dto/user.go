package dto

// RegisterInput 注册
type RegisterInput struct {
	LoginInput
	Name  string `json:"name" form:"name" comment:"昵称" example:"小明" validate:"required,alphanumunicode"` // 昵称
	Phone string `json:"phone" form:"phone" comment:"手机号" example:"19999999999" validate:"phone"`        // 手机号
}

// LoginInput 登录
type LoginInput struct {
	Input
	Email    string `json:"email" form:"email" comment:"邮箱" example:"abc123@qq.com" validate:"required,email"`       // 邮箱
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required,min=6,max=15"` // 密码
}

type LoginOutput struct {
	Id    int    `json:"id"`    // id
	Name  string `json:"name"`  // 昵称
	Email string `json:"email"` // 邮箱
}
