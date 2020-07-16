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
		}
	}

	return router
}
