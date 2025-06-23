package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func RegisterRoutes(router *gin.Engine, h *Handler) {
	fmt.Println("RegisterRoutes")
	api := router.Group("/api/user")
	{
		api.GET("/", h.GetUser)
		api.POST("/registration", h.Registration)
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	fmt.Println("GET USER CTX = ", c.Request.Context())

	fmt.Println(c.Get("userId"))
	c.JSON(http.StatusOK, "Good")
}

// Registration godoc
// @Summary Регистрация пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserTGDTO true "Информация о пользователе"
// @Success 200 {object} User
// @Router /api/user/registration [post]
func (h *Handler) Registration(c *gin.Context) {

	var req UserTGDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Println("User ID =", req.UserID)

	user := UserTGDTO{
		UserID:    req.UserID,
		UserName:  fmt.Sprintf("Test%s", req.UserID),
		FirstName: fmt.Sprintf("Test%s", req.UserID),
		LastName:  fmt.Sprintf("Test%s", req.UserID),
	}
	fmt.Println(user)
	newUser, err := h.service.RegistrationUser(user)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Good User: %s", newUser))
}
