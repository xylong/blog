package v1

import (
	"blog/internal"
	"blog/internal/dto"
	"blog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	Category = &CategoryController{}
)

type CategoryController struct {
}

// @Summary 分类列表
// @Description 分类列表
// @Tags 分类
// @Produce  json
// @Success 200 {object} dto.TagListOutput "success"
// @Router /api/v1/categories [get]
func (c *CategoryController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": service.NewCategoryService().Select(),
		"msg":  "ok",
	})
}

// @Summary 创建分类
// @Description 创建分类
// @Security ApiKeyAuth
// @Tags 分类
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.CategoryInput true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/categories [post]
func (c *CategoryController) Store(ctx *gin.Context) {
	input := &dto.CategoryInput{}
	if err := input.Check(ctx, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	if code := service.NewCategoryService().Create(input); code != internal.SUCCESS {
		internal.PanicCode(code)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 修改分类
// @Description 修改分类
// @Security ApiKeyAuth
// @Tags 分类
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.CategoryUpdate true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/categories [patch]
func (c *CategoryController) Update(ctx *gin.Context) {
	input := &dto.CategoryUpdate{}
	if err := input.Check(ctx, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	if !service.NewCategoryService().Update(input) {
		internal.PanicCode(internal.TagUpdateFail)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 删除分类
// @Description 删除分类
// @Tags 分类
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "分类ID"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/categories/{id} [delete]
func (c *CategoryController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		internal.PanicIfErr(err)
	}

	if !service.NewCategoryService().Delete(id) {
		internal.PanicCode(internal.CategoryDeleteFail)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}
