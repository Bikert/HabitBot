package router

import (
	_ "HabitMuse/docs"
	"HabitMuse/internal/middleware"
	"HabitMuse/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(spa.Middleware("/", "./webapp/dist"))
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
