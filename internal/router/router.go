package router

import (
	"HabitMuse/internal/middleware"
	"HabitMuse/internal/users"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "HabitMuse/docs"
)

func SetupRouter(userService users.Service) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/api")
	auth.Use(middleware.LogRequestBody())
	auth.Use(middleware.ValidationToken(userService))
	auth.Use(middleware.ErrorHandler())
	return router
}
