package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"

	"github.com/sirupsen/logrus"
)

type operationUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.Operation]
}

func NewOperationUsecase(repo interfaces.IRepository[model.Operation], log *logrus.Logger) interfaces.IOperationUsecase {
	return &operationUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *operationUsecase) GetAllOperations(ctx context.Context) ([]model.Operation, error) {
	u.log.Info("Fetching all operations")
	return u.Repo.GetAll(ctx)
}

func (u *operationUsecase) GetOperationByID(ctx context.Context, id uint) (*model.Operation, error) {
	u.log.Infof("Fetching operation with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *operationUsecase) CreateOperation(ctx context.Context, operation *model.Operation) error {
	u.log.Info("Creating a new operation")
	return u.Repo.Create(ctx, operation)
}

func (u *operationUsecase) UpdateOperation(ctx context.Context, operation *model.Operation) error {
	u.log.Infof("Updating operation with ID: %d", operation.ID)
	return u.Repo.Update(ctx, operation)
}

func (u *operationUsecase) DeleteOperation(ctx context.Context, id uint) error {
	u.log.Infof("Deleting operation with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
