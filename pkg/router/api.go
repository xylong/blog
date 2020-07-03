package router

import (
	v1 "blog/api/v1"
	_ "blog/docs"
	"blog/pkg/ctrl"
	"blog/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Recovery)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("ping", ctrl.Example.Ping)

	api := router.Group("/api")
	api.Use(middleware.ResponseHandler, middleware.Translate())
	{
		V1 := api.Group("/v1")
		V1.POST("/register", v1.User.Register)
		V1.POST("/login", v1.User.Login)
	}

	return router
}
