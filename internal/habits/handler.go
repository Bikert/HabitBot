package habits

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func RegisterRoutes(router *gin.Engine, h *Handler) {
	fmt.Println("RegisterRoutes")
	api := router.Group("/habits")
	{
		api.POST("/", h.CreateHabit)
		api.GET("/:date", h.GetByDate)
		api.GET("/habit/:id", h.GetHabitById)
	}
}

// CreateHabit godoc
// @Summary      Create a new habit and Update
// @Description  Creates a new habit for the authenticated user
// @Tags         habits
// @Accept       json
// @Produce      json
// @Param        habit  body      HabitDto  true  "Habit data"
// @Success      201    {object}  Habit
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /habits [post]
func (h *Handler) CreateHabit(c *gin.Context) {
	var dto *HabitDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	user := getUserByCtx(c)
	created, err := h.service.CreateHabit(dto, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create habit"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// GetHabitById godoc
// @Summary      Get habit by ID
// @Description  Retrieves a habit by its ID
// @Tags         habits
// @Produce      json
// @Param        id   path      int64  true  "Habit ID"
// @Success      200  {object}  Habit
// @Failure      404  {object}  map[string]string  "Habit not found or invalid ID"
// @Router       /habits/habit/{id} [get]
func (h *Handler) GetHabitById(c *gin.Context) {
	habitId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid habitId"})
		return
	}
	habit := h.service.GetHabitById(habitId)
	c.JSON(http.StatusOK, habit)
}

func (h *Handler) GetByDate(c *gin.Context) {

}

func getUserByCtx(c *gin.Context) *users.User {
	rawUser, exists := c.Get(constants.UserContextKey)
	if !exists {
		log.Println("user not found in context")
		return nil
	}

	user, ok := rawUser.(*users.User)
	if !ok {
		log.Println("failed to cast user from context")
		return nil
	}
	return user
}
