package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Controller() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	r.GET("/histories", GetHistoriesDetails)
	r.GET("/histories/GetHistoriesByCondition", GetHistoriesByCondition)
	r.GET("/histories/GetHistoryByCondition", GetHistoryByCondition)
	r.GET("/histories/:id", GetHistoryByID)
	r.POST("/histories", CreateHistory)
	r.PUT("/histories/:id", UpdateHistory)
	r.DELETE("/histories/:id", DeleteHistory)

	// Liquidity Pool History routes
	r.GET("/liquidityPoolHistories", GetLiquidityPoolHistoriesDetails)
	r.GET("/liquidityPoolHistories/GetLiquidityPoolHistoryByCondition", GetLiquidityPoolHistoryByCondition)

	// History Link routes
	r.GET("/historyLinks", GetHistoryLinksDetails)
	r.GET("/historyLinks/GetHistoryLinkByCondition", GetHistoryLinkByCondition)

	// Liquidity Pool History routes
	r.GET("/liquidityPoolHistoryLinks", GetLiquidityPoolHistoryLinksDetails)
	r.GET("/liquidityPoolHistoryLinks/GetLiquidityPoolHistoryLinkByCondition", GetLiquidityPoolHistoryLinkByCondition)

	return r
}
