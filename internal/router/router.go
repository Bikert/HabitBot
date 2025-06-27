package router

import (
	"HabitMuse/internal/middleware"
	"HabitMuse/internal/users"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "HabitMuse/docs"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func NewProtectedGroup(router *gin.Engine, userService users.Service) *gin.RouterGroup {
	auth := router.Group("/api")
	auth.Use(
		middleware.LogRequestBody(),
		middleware.ValidationToken(userService),
		middleware.ErrorHandler(),
	)
	return auth
}
