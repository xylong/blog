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
	Article = &ArticleController{}
)

type ArticleController struct {
}

// @Summary 文章列表
// @Description 文章列表
// @Tags 文章
// @Produce  json
// @Param category_id query int false "分类id"
// @Param title query string false "搜索关键词"
// @Param page_size query string true "条数"
// @Param page_no query string true "页码"
// @Success 200 {object} dto.ArticleListOutput "success"
// @Router /api/v1/articles [get]
func (a *ArticleController) Index(c *gin.Context) {
	input := &dto.ArticleListInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": service.NewArticleService().Select(input),
		"msg":  "ok",
	})
}

// @Summary 文章详情
// @Description 文章详情
// @Tags 文章
// @Produce  json
// @Param id path string true "文章ID"
// @Success 200 {object} dto.ArticleOutput "success"
// @Router /api/v1/articles/{id} [get]
func (a *ArticleController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		internal.PanicIfErr(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": service.NewArticleService().Find(uint(id)),
		"msg":  "ok",
	})
}

// @Summary 写文章
// @Description 写文章
// @Security ApiKeyAuth
// @Tags 文章
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.ArticleInput true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/articles [post]
func (a *ArticleController) Store(c *gin.Context) {
	input := &dto.ArticleInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}

	if code := service.NewArticleService().Create(input); code != internal.SUCCESS {
		internal.PanicCode(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 修改文章
// @Description 修改文章
// @Security ApiKeyAuth
// @Tags 文章
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param body body dto.ArticleUpdateInput true "body"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/articles [patch]
func (a *ArticleController) Update(c *gin.Context) {
	input := &dto.ArticleUpdateInput{}
	if err := input.Check(c, input); err != nil {
		internal.PanicError(http.StatusBadRequest, err)
	}
	if !service.NewArticleService().Update(input) {
		internal.PanicCode(internal.ArticleUpdateFail)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}

// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "文章ID"
// @Success 200 {json} json "{"code":200,"data":"","msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func (a *ArticleController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		internal.PanicIfErr(err)
	}

	if !service.NewArticleService().Delete(uint(id)) {
		internal.PanicCode(internal.ArticleDeleteFail)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "ok",
	})
}
