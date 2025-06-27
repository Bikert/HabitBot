package habits

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func RegisterRoutes(auth *gin.RouterGroup, h *Handler) {
	fmt.Println("RegisterRoutes")
	api := auth.Group("/habits")
	{
		api.POST("/", h.CreateHabit)
		api.GET("/:date", h.GetByDate)
		api.GET("/habit/:id", h.GetHabitById)
		api.PATCH("/habit/:id/:date", h.ToggleHabitCompletion)
	}
}

// CreateHabit godoc
// @Summary      Create a new habit and Update
// @Description  Creates a new habit for the authenticated user If id = 0 - Create new else Update
// @Tags         habits
// @Accept       json
// @Produce      json
// @Param        habit  body      HabitDto  true  "Habit data"
// @Success      201    {object}  Habit
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /api/habits [post]
func (h *Handler) CreateHabit(c *gin.Context) {
	var dto *HabitDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	user := utils.GetUserByCtx(c)
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
// @Router       /api/habits/habit/{id} [get]
func (h *Handler) GetHabitById(c *gin.Context) {
	habitId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid habitId"})
		return
	}
	habit := h.service.GetHabitById(habitId)
	c.JSON(http.StatusOK, habit)
}

// GetByDate godoc
// @Summary      Get habits by date for current user
// @Description  Returns list of habits for the authenticated user on the given date
// @Tags         habits
// @Produce      json
// @Param        date  path      string  true  "Date in YYYY-MM-DD format"  Format(date)  example(2025-06-27)
// @Success      200   {array}   Habit
// @Failure      400   {object}  map[string]string  "Invalid date format"
// @Failure      500   {object}  map[string]string  "Could not get habits"
// @Router       /api/habits/{date} [get]
func (h *Handler) GetByDate(c *gin.Context) {
	dateStr := c.Param("date")
	date, err := time.Parse(constants.DayFormat, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	user := utils.GetUserByCtx(c)
	habits, err := h.service.GetHabitsForUserByDate(user, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get habits"})
		return
	}
	c.JSON(http.StatusOK, habits)
}

// ToggleHabitCompletion godoc
// @Summary      Toggle habit completion status for a given date
// @Description  Marks a habit as completed or not completed for the specified date for the current user
// @Tags         habits
// @Accept       json
// @Produce      json
// @Param        id        path      int     true  "Habit ID"
// @Param        date      path      string  true  "Date in YYYY-MM-DD format"  Format(date)  example(2025-06-27)
// @Param        completion body      CompletionRequest true  "Completion status"
// @Success      200       {object}  map[string]bool  "Success response"
// @Failure      400       {object}  map[string]string "Invalid request data"
// @Failure      404       {object}  map[string]string "Habit not found or invalid ID"
// @Failure      500       {object}  map[string]string "Internal server error"
// @Router       /api/habits/habit/{id}/{date} [patch]
func (h *Handler) ToggleHabitCompletion(c *gin.Context) {
	habitId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid habitId"})
		return
	}
	dateStr := c.Param("date")
	date, err := time.Parse(constants.DayFormat, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	var req CompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	completed := req.Completed

	user := utils.GetUserByCtx(c)
	err = h.service.ToggleHabitCompletion(user, habitId, date, completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
