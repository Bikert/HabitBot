package router

import (
	"HabitMuse/internal/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "HabitMuse/docs"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(http.LogRequestBody())
	//router.Use(http.ValidationToken())
	router.Use(http.ErrorHandler())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
