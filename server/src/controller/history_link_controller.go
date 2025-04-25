package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHistoryLinks(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	historyLinks, err := module.GetAllHistoryLinks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historyLinks)
}

// @Summary Get detailed historyLinks
// @Description Retrieve detailed historyLink records with related entities
// @Tags HistoryLinks
// @Accept json
// @Produce json
// @Success 200 {array} dto.HistoryLinkResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /historyLinks [get]
func GetHistoryLinksDetails(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	historyLinks, err := module.GetHistoryLinksDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historyLinks)
}

// @Summary Get historyLink by condition
// @Description Retrieve historyLink based on filter conditions
// @Tags HistoryLinks
// @Accept json
// @Produce json
// @Param provider query string false "Provider name"
// @Param token query string false "Token name"
// @Param operation query string false "Operation name"
// @Success 200 {object} dto.HistoryLinkResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /historyLinks/GetHistoryLinkByCondition [get]
func GetHistoryLinkByCondition(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	var req dto.GetHistoryLinkRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	historyLinks, err := module.GetHistoryLinkByCondition(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historyLinks)
}

func GetHistoryLinkByID(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	historyLink, err := module.GetHistoryLinkByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historyLink)
}

func CreateHistoryLink(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	var historyLink model.HistoryLink
	if err := ctx.ShouldBindJSON(&historyLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.CreateHistoryLink(ctx, &historyLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, historyLink)
}

func UpdateHistoryLink(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	var historyLink model.HistoryLink
	if err := ctx.ShouldBindJSON(&historyLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.UpdateHistoryLink(ctx, &historyLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, historyLink)
}

func DeleteHistoryLink(ctx *gin.Context) {
	module := dependency_injection.NewHistoryLinkUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	err = module.DeleteHistoryLink(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
