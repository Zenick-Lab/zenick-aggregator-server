package interfaces

import (
	"context"

	"gorm.io/gorm"
)

type IRepository[T any] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id uint) (*T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
	GetDB() *gorm.DB
}
