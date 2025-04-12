package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"
)

type IOperationUsecase interface {
	GetAllOperations(ctx context.Context) ([]model.Operation, error)
	GetOperationByID(ctx context.Context, id uint) (*model.Operation, error)
	CreateOperation(ctx context.Context, operation *model.Operation) error
	UpdateOperation(ctx context.Context, operation *model.Operation) error
	DeleteOperation(ctx context.Context, id uint) error
}
