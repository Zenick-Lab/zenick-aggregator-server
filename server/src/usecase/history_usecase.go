package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"

	"github.com/sirupsen/logrus"
)

type historyUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.History]
}

func NewHistoryUsecase(repo interfaces.IRepository[model.History], log *logrus.Logger) interfaces.IHistoryUsecase {
	return &historyUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *historyUsecase) GetAllHistories(ctx context.Context) ([]model.History, error) {
	u.log.Info("Fetching all histories")
	return u.Repo.GetAll(ctx)
}

func (u *historyUsecase) GetHistoryByID(ctx context.Context, id uint) (*model.History, error) {
	u.log.Infof("Fetching history with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *historyUsecase) CreateHistory(ctx context.Context, history *model.History) error {
	u.log.Info("Creating a new history")
	return u.Repo.Create(ctx, history)
}

func (u *historyUsecase) UpdateHistory(ctx context.Context, history *model.History) error {
	u.log.Infof("Updating history with ID: %d", history.ID)
	return u.Repo.Update(ctx, history)
}

func (u *historyUsecase) DeleteHistory(ctx context.Context, id uint) error {
	u.log.Infof("Deleting history with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
