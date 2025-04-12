package controller

import (
	"Zenick-Lab/zenick-aggregator-server/src/dependency_injection"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllOperations(ctx *gin.Context) {
	module := dependency_injection.NewOperationUsecaseProvider()

	operations, err := module.GetAllOperations(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, operations)
}

func GetOperationByID(ctx *gin.Context) {
	module := dependency_injection.NewOperationUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	operation, err := module.GetOperationByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func CreateOperation(ctx *gin.Context) {
	module := dependency_injection.NewOperationUsecaseProvider()

	var operation model.Operation
	if err := ctx.ShouldBindJSON(&operation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.CreateOperation(ctx, &operation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, operation)
}

func UpdateOperation(ctx *gin.Context) {
	module := dependency_injection.NewOperationUsecaseProvider()

	var operation model.Operation
	if err := ctx.ShouldBindJSON(&operation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := module.UpdateOperation(ctx, &operation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func DeleteOperation(ctx *gin.Context) {
	module := dependency_injection.NewOperationUsecaseProvider()

	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	id := uint(idUint)

	err = module.DeleteOperation(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
