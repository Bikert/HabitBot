package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(router gin.IRouter) {
	fmt.Println("RegisterRoutes")

	api := router.Group("/users")

	api.GET("/", h.GetUser)
}

// GetUser godoc
// @Summary Получить Пользователя из контекста
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /api/user/ [get]
func (h *Handler) GetUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "message: User Not Found")
	}
	fmt.Println("USER = ", user)
	c.JSON(http.StatusOK, user)
}
