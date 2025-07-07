package body_metrics

import (
	"HabitMuse/internal/dto"
	"HabitMuse/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router gin.IRouter) {
	api := router.Group("/body-metric")
	api.GET("/all", h.GetAllMetric)
	api.GET("/report", h.GetReport)
	api.POST("/create", h.Create)
}

// Create godoc
// @Summary Создать метрику тела
// @Tags BodyMetrics
// @Accept json
// @Produce json
// @Param metric body BodyMetricDTO true "Данные метрики"
// @Success 201 {object} BodyMetricDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/body-metric/create [post]
func (h *Handler) Create(ctx *gin.Context) {
	user := utils.GetUserByCtx(ctx)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "user not found in context",
		})
		return
	}
	var bodyMetricDTO *BodyMetricDTO
	if err := ctx.ShouldBindJSON(&bodyMetricDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	log.Println("BodyMetricDTO = ", bodyMetricDTO)

	newDto, err := h.service.Create(user.UserID, bodyMetricDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newDto)
}

// GetReport godoc
// @Summary Получить отчет по метрикам тела
// @Tags BodyMetrics
// @Produce json
// @Success 200 {string} string "Готовый текстовый отчет"
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/body-metric [get]
func (h *Handler) GetReport(ctx *gin.Context) {
	user := utils.GetUserByCtx(ctx)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "user not found in context",
		})
		return
	}

	str, err := h.service.GetReport(user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}
	ctx.JSON(http.StatusOK, str)
}

// GetAllMetric godoc
// @Summary      Получить список всех метрик тела пользователя
// @Description  Возвращает список всех записей метрик тела текущего авторизованного пользователя
// @Tags         body-metric
// @Produce      json
// @Success      200 {array} BodyMetricDTO
// @Failure      401 {object} dto.ErrorResponse "Пользователь не найден в контексте"
// @Failure      500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /body-metric/all [get]
func (h *Handler) GetAllMetric(ctx *gin.Context) {
	user := utils.GetUserByCtx(ctx)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "user not found in context",
		})
		return
	}

	list, err := h.service.GetMetricsList(user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
