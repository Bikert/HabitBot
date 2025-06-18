package habits

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

func RegisterRoutes(router *gin.Engine, h *Handler) {
	fmt.Println("RegisterRoutes")
	api := router.Group("/api")
	{
		api.GET("/habits", h.GetHabits)
	}
}

func (h *Handler) GetHabits(c *gin.Context) {
	fmt.Println("GET HABITS CTX = ", c.Request.Context())
	//userId := c.MustGet("user_id").(int64)
	//habits, err := h.service.GetHabitsByUser(userId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get habits"})
	//	return
	//}
	//c.JSON(http.StatusOK, habits)
	c.JSON(http.StatusOK, "Good")
}

func (h *Handler) CreateHabit(c *gin.Context) {
	var habit Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	created, err := h.service.CreateNewHabit(habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create habit"})
		return
	}
	c.JSON(http.StatusCreated, created)
}
