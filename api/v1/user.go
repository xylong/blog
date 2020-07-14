package v1

import (
	"blog/pkg"
	"blog/pkg/dto"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var User = &UserController{}

//  UserController 用户控制器
type UserController struct {
}

// @Summary 注册
// @Description 用户注册
// @Tags 用户接口
// @Produce  json
// @Param body body dto.RegisterInput true "body"
// @Success 200 {object} dto.LoginOutput "success"
// @Router /api/v1/register [post]
func (u *UserController) Register(c *gin.Context) {
	input := &dto.RegisterInput{}
	if err := input.Check(c, input); err != nil {
		pkg.PanicError(http.StatusBadRequest, err)
	}

	serve := service.NewUserService()
	if err := serve.Register(input); err != nil {
		pkg.PanicIfErr(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 登录
// @Description 用户登录
// @Tags 用户接口
// @Produce  json
// @Param body body dto.LoginInput true "body"
// @Success 200 {object} dto.LoginOutput "success"
// @Router /api/v1/login [post]
func (u *UserController) Login(c *gin.Context) {
	input := &dto.LoginInput{}
	if err := input.Check(c, input); err != nil {
		pkg.PanicError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": input.Email,
		"msg":  "ok",
	})
}
