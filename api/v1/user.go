package v1

import (
	"blog/internal"
	"blog/internal/dto"
	"blog/internal/service"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	User = &UserController{}
)

//  UserController 用户控制器
type UserController struct {
}

// @Summary 注册
// @Description 用户注册
// @Tags 用户接口
// @Produce  json
// @Param body body dto.RegisterInput true "body"
// @Success 200 {json} json "{"code":200,"data":null,"msg":"ok"}"
// @Router /api/v1/register [post]
func (u *UserController) Register(c *gin.Context) {
	input := &dto.RegisterInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	serve := service.NewUserService()
	if err := serve.Register(input); err != nil {
		internal.PanicIfErr(err)
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
// @Success 200 {json} json "{"code":200,"data":"token","msg":"ok"}"
// @Router /api/v1/login [post]
func (u *UserController) Login(c *gin.Context) {
	input := &dto.LoginInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}

	serve := service.NewUserService()
	token, err := serve.Login(input)
	if err != nil {
		internal.PanicIfErr(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": token,
		"msg":  "ok",
	})
}

// @Summary 个人信息
// @Description 获取个人信息
// @Tags 用户接口
// @Security ApiKeyAuth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} dto.Profile "success"
// @Router /api/v1/user [get]
func (u *UserController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*util.Claims)
	id, _ := strconv.Atoi(claims.ID)
	serve := service.NewUserService()
	user, err := serve.Profile(uint(id))
	if err != nil {
		internal.PanicIfErr(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": user,
		"msg":  "ok",
	})
}
