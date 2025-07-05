package abstractions

import "github.com/gin-gonic/gin"

type ApiHandler interface {
	RegisterRoutes(router gin.IRouter)
}
