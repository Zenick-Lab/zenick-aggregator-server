package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"
)

type ILiquidityPoolHistoryUsecase interface {
	GetAllLiquidityPoolHistories(ctx context.Context) ([]model.LiquidityPoolHistory, error)
	GetLiquidityPoolHistoriesDetails(ctx context.Context) ([]dto.LiquidityPoolHistoryResponse, error)
	GetLiquidityPoolHistoryByCondition(ctx context.Context, req *dto.GetNewestLiquidityPoolHistoryRequest) (*dto.LiquidityPoolHistoryResponse, error)
	GetLiquidityPoolHistoryByID(ctx context.Context, id uint) (*model.LiquidityPoolHistory, error)
	CreateLiquidityPoolHistory(ctx context.Context, history *model.LiquidityPoolHistory) error
	UpdateLiquidityPoolHistory(ctx context.Context, history *model.LiquidityPoolHistory) error
	DeleteLiquidityPoolHistory(ctx context.Context, id uint) error
}
