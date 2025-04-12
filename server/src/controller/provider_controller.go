package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProviders(ctx *gin.Context) {
	module := dependency_injection.NewProviderUsecaseProvider()

	providers, err := module.GetAllProviders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, providers)
}

func GetProviderByID(ctx *gin.Context) {
	module := dependency_injection.NewProviderUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	provider, err := module.GetProviderByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

func CreateProvider(ctx *gin.Context) {
	module := dependency_injection.NewProviderUsecaseProvider()

	var provider model.Provider
	if err := ctx.ShouldBindJSON(&provider); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.CreateProvider(ctx, &provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, provider)
}

func UpdateProvider(ctx *gin.Context) {
	module := dependency_injection.NewProviderUsecaseProvider()

	var provider model.Provider
	if err := ctx.ShouldBindJSON(&provider); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.UpdateProvider(ctx, &provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

func DeleteProvider(ctx *gin.Context) {
	module := dependency_injection.NewProviderUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	err = module.DeleteProvider(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
