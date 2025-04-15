package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHistories(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	histories, err := module.GetAllHistories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, histories)
}

// @Summary Get detailed histories
// @Description Retrieve detailed history records with related entities
// @Tags Histories
// @Accept json
// @Produce json
// @Success 200 {array} dto.HistoryResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /histories [get]
func GetHistoriesDetails(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	histories, err := module.GetHistoriesDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, histories)
}

// @Summary Get histories by condition
// @Description Retrieve histories based on filter conditions
// @Tags Histories
// @Accept json
// @Produce json
// @Param provider query string false "Provider name"
// @Param token query string false "Token name"
// @Param operation query string false "Operation name"
// @Param apr query number false "APR value"
// @Param from_date query string false "Start date"
// @Param to_date query string false "End date"
// @Success 200 {array} dto.HistoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /histories/GetHistoriesByCondition [get]
func GetHistoriesByCondition(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	var req dto.GetHistoryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	histories, err := module.GetHistoriesByCondition(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, histories)
}

// @Summary Get history by ID
// @Description Retrieve a single history record by its ID
// @Tags Histories
// @Accept json
// @Produce json
// @Param id path int true "History ID"
// @Success 200 {object} model.History
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /histories/{id} [get]
func GetHistoryByID(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	history, err := module.GetHistoryByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func CreateHistory(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	var history model.History
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.CreateHistory(ctx, &history)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, history)
}

func UpdateHistory(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	var history model.History
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.UpdateHistory(ctx, &history)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func DeleteHistory(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	err = module.DeleteHistory(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
