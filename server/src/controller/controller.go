package controller

import "github.com/gin-gonic/gin"

func Controller() *gin.Engine {
	r := gin.Default()

	// Token routes
	r.GET("/tokens", GetAllTokens)
	r.GET("/tokens/:id", GetTokenByID)
	r.POST("/tokens", CreateToken)
	r.PUT("/tokens/:id", UpdateToken)
	r.DELETE("/tokens/:id", DeleteToken)

	// Provider routes
	r.GET("/providers", GetAllProviders)
	r.GET("/providers/:id", GetProviderByID)
	r.POST("/providers", CreateProvider)
	r.PUT("/providers/:id", UpdateProvider)
	r.DELETE("/providers/:id", DeleteProvider)

	// Operation routes
	r.GET("/operations", GetAllOperations)
	r.GET("/operations/:id", GetOperationByID)
	r.POST("/operations", CreateOperation)
	r.PUT("/operations/:id", UpdateOperation)
	r.DELETE("/operations/:id", DeleteOperation)

	// History routes
	r.GET("/histories", GetAllHistories)
	r.GET("/histories/:id", GetHistoryByID)
	r.POST("/histories", CreateHistory)
	r.PUT("/histories/:id", UpdateHistory)
	r.DELETE("/histories/:id", DeleteHistory)

	return r
}
