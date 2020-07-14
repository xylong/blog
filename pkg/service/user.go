package service

import (
	"blog/pkg/dao"
	"blog/pkg/dto"
	"blog/pkg/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// 注册
	Register(*dto.RegisterInput) error
	// 登录
	Login(*dto.LoginInput) (string, error)
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

	d := dao.NewDao()
	if d.IsExist(input.Email, input.Phone) {
		return errors.New("已注册")
	}

	_, err = d.Create(&model.User{
		Name:     input.Name,
		Phone:    input.Phone,
		Email:    input.Email,
		Password: string(hash),
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Login(input *dto.LoginInput) (token string, err error) {

}

func (u *user) Logout() {

}
