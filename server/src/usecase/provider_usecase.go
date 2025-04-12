package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"

	"github.com/sirupsen/logrus"
)

type providerUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.Provider]
}

func NewProviderUsecase(repo interfaces.IRepository[model.Provider], log *logrus.Logger) interfaces.IProviderUsecase {
	return &providerUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *providerUsecase) GetAllProviders(ctx context.Context) ([]model.Provider, error) {
	u.log.Info("Fetching all providers")
	return u.Repo.GetAll(ctx)
}

func (u *providerUsecase) GetProviderByID(ctx context.Context, id uint) (*model.Provider, error) {
	u.log.Infof("Fetching provider with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *providerUsecase) CreateProvider(ctx context.Context, provider *model.Provider) error {
	u.log.Info("Creating a new provider")
	return u.Repo.Create(ctx, provider)
}

func (u *providerUsecase) UpdateProvider(ctx context.Context, provider *model.Provider) error {
	u.log.Infof("Updating provider with ID: %d", provider.ID)
	return u.Repo.Update(ctx, provider)
}

func (u *providerUsecase) DeleteProvider(ctx context.Context, id uint) error {
	u.log.Infof("Deleting provider with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
