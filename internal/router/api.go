package router

import (
	v1 "blog/api/v1"
	_ "blog/docs"
	"blog/internal/ctrl"
	"blog/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.MaxMultipartMemory = 1024
	router.Use(middleware.Recovery, middleware.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("ping", ctrl.Example.Ping)

	api := router.Group("/api")
	api.Use(middleware.ResponseHandler, middleware.Translate())
	{
		V1 := api.Group("/v1")
		// 注册
		V1.POST("/register", v1.User.Register)
		// 登录
		V1.POST("/login", v1.User.Login)
		// 标签
		V1.GET("/tags", v1.Tag.Index)
		// 分类
		V1.GET("/categories", v1.Category.Index)
		// 文章
		V1.GET("/articles", v1.Article.Index)
		V1.GET("/articles/:id", v1.Article.Show)

		V1.Use(middleware.JWT())
		{
			// 个人信息
			V1.GET("/user", v1.User.Me)
			// 标签
			tags := V1.Group("/tags")
			{
				tags.POST("", v1.Tag.Store)
				tags.PUT("", v1.Tag.Update)
				tags.DELETE(":id", v1.Tag.Delete)
			}
			// 分类
			categories := V1.Group("/categories")
			{
				categories.POST("", v1.Category.Store)
				categories.PATCH("", v1.Category.Update)
				categories.DELETE(":id", v1.Category.Delete)
			}
			// 文章
			articles := V1.Group("/articles")
			{
				articles.POST("", v1.Article.Store)
				articles.DELETE(":id", v1.Article.Delete)
				articles.PATCH("", v1.Article.Update)
			}
			// 上传
			V1.POST("/upload", v1.Upload.Upload)
		}
	}

	return router
}
