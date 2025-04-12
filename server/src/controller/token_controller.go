package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTokens(ctx *gin.Context) {
	module := dependency_injection.NewTokenUsecaseProvider()

	tokens, err := module.GetAllTokens(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tokens)
}

func GetTokenByID(ctx *gin.Context) {
	module := dependency_injection.NewTokenUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	token, err := module.GetTokenByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func CreateToken(ctx *gin.Context) {
	module := dependency_injection.NewTokenUsecaseProvider()

	var token model.Token
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.CreateToken(ctx, &token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, token)
}

func UpdateToken(ctx *gin.Context) {
	module := dependency_injection.NewTokenUsecaseProvider()

	var token model.Token
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.UpdateToken(ctx, &token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func DeleteToken(ctx *gin.Context) {
	module := dependency_injection.NewTokenUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	err = module.DeleteToken(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
