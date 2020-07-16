package service

import (
	"blog/internal/dao"
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/util"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserService interface {
	// 注册
	Register(*dto.RegisterInput) error
	// 登录
	Login(*dto.LoginInput) (string, error)
	// 注销
	Logout()
	// 个人信息
	Profile(uint) (*dto.Profile, error)
}

func NewUserService() UserService {
	return &user{
		dao: dao.NewUserDao(),
	}
}

type user struct {
	dao dao.UserDao
}

func (u *user) Register(input *dto.RegisterInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if u.dao.IsExist(input.Email, input.Phone) {
		return errors.New("已注册")
	}

	_, err = u.dao.Create(&model.User{
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
	info, err := u.dao.Find(&model.User{
		Email: input.Email,
	})
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("密码错误")
	}
	id := strconv.Itoa(int(info.ID))
	token, err = util.NewJWT().Generate(&util.Claims{
		ID:    id,
		Name:  info.Name,
		Email: info.Email,
	})
	return
}

func (u *user) Logout() {

}

func (u *user) Profile(id uint) (*dto.Profile, error) {
	me := &model.User{}
	me.ID = id
	me, err := u.dao.Find(me)
	if err != nil {
		return nil, err
	}
	profile := &dto.Profile{
		Id:     me.ID,
		Avatar: me.Avatar,
		Name:   me.Name,
		Email:  me.Email,
		Phone:  me.Phone,
	}
	return profile, nil
}
