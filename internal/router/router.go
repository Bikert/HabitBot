package router

import (
	"HabitMuse/internal/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(http.LogRequestBody())
	router.Use(http.ErrorHandler())
	return router
}
