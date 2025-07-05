package http

import (
	_ "HabitMuse/docs"
	"HabitMuse/internal/abstractions"
	"HabitMuse/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewEngine(apiHandlers []abstractions.ApiHandler, userService users.Service) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := router.Group("/api")
	apiGroup.Use(
		LogRequestBody(),
		ValidationToken(userService),
		ErrorHandler(),
	)

	for _, apiHandler := range apiHandlers {
		apiHandler.RegisterRoutes(apiGroup)
	}

	// HACK: spa middleware MUST be registered after all other middlewares
	// TODO: reverse router registration logic along with moving it to the single Provide phase
	router.Use(spa.Middleware("/", "./webapp/dist"))
	return router
}
