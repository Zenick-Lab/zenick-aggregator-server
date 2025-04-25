package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get detailed liquidityPoolHistoryLinks
// @Description Retrieve detailed liquidityPoolHistory records with related entities
// @Tags LiquidityPoolHistoryLinks
// @Accept json
// @Produce json
// @Success 200 {array} dto.LiquidityPoolHistoryLinkResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /liquidityPoolHistoryLinks [get]
func GetLiquidityPoolHistoryLinksDetails(ctx *gin.Context) {
	module := dependency_injection.NewLiquidityPoolHistoryLinkUsecaseProvider()

	liquidityPoolHistories, err := module.GetLiquidityPoolHistoryLinksDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, liquidityPoolHistories)
}

// @Summary Get Liquidity Pool History Link by condition
// @Description Retrieve Liquidity Pool History Link based on filter conditions
// @Tags LiquidityPoolHistoryLinks
// @Accept json
// @Produce json
// @Param provider query string false "Provider name"
// @Param token_a query string false "Token A name"
// @Param token_b query string false "Token B name"
// @Success 200 {object} dto.LiquidityPoolHistoryLinkResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /liquidityPoolHistoryLinks/GetLiquidityPoolHistoryLinkByCondition [get]
func GetLiquidityPoolHistoryLinkByCondition(ctx *gin.Context) {
	module := dependency_injection.NewLiquidityPoolHistoryLinkUsecaseProvider()

	var req dto.GetLiquidityPoolHistoryLinkRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	liquidityPoolHistories, err := module.GetLiquidityPoolHistoryLinkByCondition(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, liquidityPoolHistories)
}
