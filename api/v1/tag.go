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
	Tag = &TagController{}
)

type TagController struct {
}

// @Summary 标签列表
// @Description 标签列表
// @Tags 标签
// @Produce  json
// @Success 200 {object} dto.TagListOutput "success"
// @Router /api/v1/tags [get]
func (t *TagController) Index(c *gin.Context) {
	serve := service.NewTagService()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": serve.Select(),
		"msg":  "ok",
	})
}

// @Summary 创建标签
// @Description 创建标签
// @Security ApiKeyAuth
// @Tags 标签
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.TagInput true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/tags [post]
func (t *TagController) Store(c *gin.Context) {
	input := &dto.TagInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	serve := service.NewTagService()
	if code := serve.Create(input); code != internal.SUCCESS {
		internal.PanicCode(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 修改标签
// @Description 修改标签
// @Security ApiKeyAuth
// @Tags 标签
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.TagUpdate true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/tags [put]
func (t *TagController) Update(c *gin.Context) {
	input := &dto.TagUpdate{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	serve := service.NewTagService()
	if !serve.Update(input) {
		internal.PanicCode(internal.TagUpdateFail)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 删除标签
// @Description 删除标签
// @Tags 标签
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "标签ID"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func (t *TagController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		internal.PanicIfErr(err)
	}

	if !service.NewTagService().Delete(id) {
		internal.PanicCode(internal.TagDeleteFail)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}
