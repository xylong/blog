package v1

import (
	"github.com/gin-gonic/gin"
)

var (
	Article = &ArticleController{}
)

type ArticleController struct {
}

func (a *ArticleController) Index(c *gin.Context) {

}

func (a *ArticleController) Show(c *gin.Context) {

}

func (a *ArticleController) Store(c *gin.Context) {

}

func (a *ArticleController) Update(c *gin.Context) {

}

func (a *ArticleController) Delete(c *gin.Context) {

}
