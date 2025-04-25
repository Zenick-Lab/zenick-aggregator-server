package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"
)

type IHistoryUsecase interface {
	GetAllHistories(ctx context.Context) ([]model.History, error)
	GetHistoriesDetails(ctx context.Context) ([]dto.HistoryResponse, error)
	GetHistoriesByCondition(ctx context.Context, req *dto.GetNewestHistoryRequest) ([]dto.HistoryResponse, error)
	GetHistoryByCondition(ctx context.Context, req *dto.GetNewestHistoryRequest) (*dto.HistoryResponse, error)
	GetHistoryByID(ctx context.Context, id uint) (*model.History, error)
	CreateHistory(ctx context.Context, history *model.History) error
	UpdateHistory(ctx context.Context, history *model.History) error
	DeleteHistory(ctx context.Context, id uint) error
}
