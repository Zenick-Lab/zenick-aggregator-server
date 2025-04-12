package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"
)

type IHistoryUsecase interface {
	GetAllHistories(ctx context.Context) ([]model.History, error)
	GetHistoryByID(ctx context.Context, id uint) (*model.History, error)
	CreateHistory(ctx context.Context, history *model.History) error
	UpdateHistory(ctx context.Context, history *model.History) error
	DeleteHistory(ctx context.Context, id uint) error
}
