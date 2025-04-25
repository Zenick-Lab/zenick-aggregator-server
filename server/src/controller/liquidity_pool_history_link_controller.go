package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLiquidityPoolHistoryLinksDetails(ctx *gin.Context) {
	module := dependency_injection.NewLiquidityPoolHistoryLinkUsecaseProvider()

	liquidityPoolHistories, err := module.GetLiquidityPoolHistoryLinksDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, liquidityPoolHistories)
}

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
