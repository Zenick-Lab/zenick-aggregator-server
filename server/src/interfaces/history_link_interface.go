package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"
)

type IHistoryLinkUsecase interface {
	GetAllHistoryLinks(ctx context.Context) ([]model.HistoryLink, error)
	GetHistoryLinksDetails(ctx context.Context) ([]dto.HistoryLinkResponse, error)
	GetHistoryLinkByCondition(ctx context.Context, req *dto.GetHistoryLinkRequest) (*dto.HistoryLinkResponse, error)
	GetHistoryLinkByID(ctx context.Context, id uint) (*model.HistoryLink, error)
	CreateHistoryLink(ctx context.Context, history *model.HistoryLink) error
	UpdateHistoryLink(ctx context.Context, history *model.HistoryLink) error
	DeleteHistoryLink(ctx context.Context, id uint) error
}
