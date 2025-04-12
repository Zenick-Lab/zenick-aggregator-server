package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHistories(ctx *gin.Context) {
	module := dependency_injection.NewHistoryUsecaseProvider()

	historys, err := module.GetAllHistories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historys)
}

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
