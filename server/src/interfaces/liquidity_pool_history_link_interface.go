package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"
)

type ILiquidityPoolHistoryLinkUsecase interface {
	GetAllLiquidityPoolHistoryLinks(ctx context.Context) ([]model.LiquidityPoolHistoryLink, error)
	GetLiquidityPoolHistoryLinksDetails(ctx context.Context) ([]dto.LiquidityPoolHistoryLinkResponse, error)
	GetLiquidityPoolHistoryLinkByCondition(ctx context.Context, req *dto.GetLiquidityPoolHistoryLinkRequest) (*dto.LiquidityPoolHistoryLinkResponse, error)
	GetLiquidityPoolHistoryLinkByID(ctx context.Context, id uint) (*model.LiquidityPoolHistoryLink, error)
	CreateLiquidityPoolHistoryLink(ctx context.Context, history *model.LiquidityPoolHistoryLink) error
	UpdateLiquidityPoolHistoryLink(ctx context.Context, history *model.LiquidityPoolHistoryLink) error
	DeleteLiquidityPoolHistoryLink(ctx context.Context, id uint) error
}
