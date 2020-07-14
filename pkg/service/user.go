package service

import (
	"blog/pkg/dao"
	"blog/pkg/dto"
	"blog/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// 注册
	Register(input *dto.RegisterInput) error
	// 登录
	Login()
	// 注销
	Logout()
}

func NewUserService() UserService {
	return &user{}
}

type user struct {
}

func (u *user) Register(input *dto.RegisterInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = dao.NewDao().Create(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: string(hash),
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Login() {

}

func (u *user) Logout() {

}
