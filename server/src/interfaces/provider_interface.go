package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"
)

type IProviderUsecase interface {
	GetAllProviders(ctx context.Context) ([]model.Provider, error)
	GetProviderByID(ctx context.Context, id uint) (*model.Provider, error)
	CreateProvider(ctx context.Context, provider *model.Provider) error
	UpdateProvider(ctx context.Context, provider *model.Provider) error
	DeleteProvider(ctx context.Context, id uint) error
}
