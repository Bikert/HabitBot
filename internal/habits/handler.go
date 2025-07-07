package habits

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/dto"
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

func (h *Handler) RegisterRoutes(auth gin.IRouter) {
	fmt.Println("RegisterRoutes")
	api := auth.Group("/habit")
	{
		api.GET("/", h.getAllActiveHabits)
		api.GET("/:groupId", h.GetActiveHabitByGroupId)
		api.GET("/completion/:date", h.GetCompletionHabitsByDate)
		api.POST("/create", h.CreateHabit)
		api.PUT("/update/:groupId", h.Update)
		api.PATCH("/:versionId/:date", h.ToggleHabitCompletion)
	}
}

// getAllActiveHabits @Summary Get all active habits
// @Description Returns all active habits for the authenticated user
// @Tags habits
// @Produce json
// @Success 200 {array} HabitDto
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/habit/ [get]
func (h *Handler) getAllActiveHabits(c *gin.Context) {
	user := utils.GetUserByCtx(c)
	habitDto, err := h.service.GetHabitsByUser(user.UserID)
	log.Println("getAllActiveHabits", habitDto, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habitDto)
}

// GetActiveHabitByGroupId @Summary Get active habit by group ID
// @Description Returns a single active habit by group ID
// @Tags habits
// @Produce json
// @Param groupId path string true "Habit group ID"
// @Success 200 {object} HabitDto
// @Router /api/habit/{groupId} [get]
func (h *Handler) GetActiveHabitByGroupId(c *gin.Context) {
	groupId := c.Param("groupId")
	habit := h.service.GetHabitByGroupID(groupId)
	c.JSON(http.StatusOK, habit)
}

// GetCompletionHabitsByDate @Summary Get completed habits for a specific date
// @Description Returns all completed or active habits for a user on a given date
// @Tags habits
// @Produce json
// @Param date path string true "Date in format YYYY-MM-DD"
// @Success 200 {array} HabitCompletionDto
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/habit/completion/{date} [get]
func (h *Handler) GetCompletionHabitsByDate(c *gin.Context) {
	dateStr := c.Param("date")
	date, err := time.Parse(constants.DayFormat, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf("Invalid date format: %s", dateStr)})
		return
	}
	user := utils.GetUserByCtx(c)
	habits, err := h.service.GetCompletionHabitsForUserByDate(user, date)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get habits"})
		return
	}
	c.JSON(http.StatusOK, habits)
}

// CreateHabit godoc
// @Summary Создать новую привычку
// @Tags habits
// @Accept json
// @Produce json
// @Param request body CreateHabitDto true "CreateHabitDto"
// @Success 201 {object} HabitDto
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/habit/create [post]
func (h *Handler) CreateHabit(c *gin.Context) {
	var habitDto *CreateHabitDto
	log.Println("CreateHabit")
	if err := c.ShouldBindJSON(&habitDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("CreateHabit = ", habitDto)
	user := utils.GetUserByCtx(c)
	created, err := h.service.CreateHabit(habitDto, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Update godoc
// @Summary Обновить привычку
// @Tags habits
// @Accept json
// @Produce json
// @Param groupId path string true "ID группы привычки"
// @Param request body UpdateHabitDto true "UpdateHabitDto"
// @Success 200 {object} HabitDto
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/habit/update/{groupId} [put]
func (h *Handler) Update(c *gin.Context) {
	var habitDto *UpdateHabitDto
	if err := c.ShouldBindJSON(&habitDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupId := c.Param("groupId")
	user := utils.GetUserByCtx(c)
	updated, err := h.service.UpdateHabit(groupId, habitDto, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update habit"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// ToggleHabitCompletion godoc
// @Summary Отметить выполнение привычки на дату
// @Tags habits
// @Accept json
// @Produce json
// @Param versionId path int true "ID версии привычки"
// @Param date path string true "Дата (в формате 2006-01-02)"
// @Param request body CompletionRequest true "Completion status"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/habit/{versionId}/{date} [patch]
func (h *Handler) ToggleHabitCompletion(c *gin.Context) {
	habitId, err := strconv.ParseInt(c.Param("versionId"), 10, 64)
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
		log.Printf("Error binding request for  habitId=%v, date=%v: %v", habitId, dateStr, err)
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
