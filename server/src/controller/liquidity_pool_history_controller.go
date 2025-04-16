package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get detailed liquidityPoolHistories
// @Description Retrieve detailed liquidityPoolHistory records with related entities
// @Tags LiquidityPoolHistories
// @Accept json
// @Produce json
// @Success 200 {array} dto.LiquidityPoolHistoryResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /liquidityPoolHistories [get]
func GetLiquidityPoolHistoriesDetails(ctx *gin.Context) {
	module := dependency_injection.NewLiquidityPoolHistoryUsecaseProvider()

	liquidityPoolHistories, err := module.GetLiquidityPoolHistoriesDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, liquidityPoolHistories)
}

// @Summary Get Liquidity Pool History by condition
// @Description Retrieve Liquidity Pool History based on filter conditions
// @Tags LiquidityPoolHistories
// @Accept json
// @Produce json
// @Param provider query string false "Provider name"
// @Success 200 {object} dto.LiquidityPoolHistoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /liquidityPoolHistories/GetLiquidityPoolHistoryByCondition [get]
func GetLiquidityPoolHistoryByCondition(ctx *gin.Context) {
	module := dependency_injection.NewLiquidityPoolHistoryUsecaseProvider()

	var req dto.GetNewestLiquidityPoolHistoryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	liquidityPoolHistories, err := module.GetLiquidityPoolHistoryByCondition(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, liquidityPoolHistories)
}
