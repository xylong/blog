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

		V1.Use(middleware.JWT())
		{
			// 个人信息
			V1.GET("/user", v1.User.Me)
			// 标签
			V1.GET("/tags", v1.Tag.Index)
			V1.GET("/tags/:id", v1.Tag.Show)
			V1.POST("/tags", v1.Tag.Store)
			V1.PUT("/tags/:id", v1.Tag.Update)
			V1.DELETE("/tags/:id", v1.Tag.Delete)
			// 分类
			V1.GET("/categories", v1.Category.Index)
			V1.GET("/categories/:id", v1.Category.Show)
			V1.POST("/categories", v1.Category.Store)
			V1.PUT("/categories/:id", v1.Category.Update)
			V1.DELETE("/categories/:id", v1.Category.Delete)
			// 文章
			V1.GET("/articles", v1.Article.Index)
			V1.GET("/articles/:id", v1.Article.Show)
			V1.POST("/articles", v1.Article.Store)
			V1.PUT("/articles/:id", v1.Article.Update)
			V1.DELETE("/articles/:id", v1.Article.Delete)
		}
	}

	return router
}
