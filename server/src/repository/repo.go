package repository

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repository[T any] struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewRepository[T any](db *gorm.DB, log *logrus.Logger) interfaces.IRepository[T] {
	return &repository[T]{
		db:  db,
		log: log,
	}
}

func (r *repository[T]) GetAll(ctx context.Context) ([]T, error) {
	r.log.Info("Fetching all entities")

	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		r.log.Errorf("Error fetching all entities: %v", err)
	}
	return entities, err
}

func (r *repository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	r.log.Infof("Fetching entity with ID: %d", id)

	var entity T
	err := r.db.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		r.log.Errorf("Error fetching entity with ID %d: %v", id, err)
	}
	return &entity, err
}

func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	r.log.Info("Creating a new entity")

	err := r.db.WithContext(ctx).Create(entity).Error
	if err != nil {
		r.log.Errorf("Error creating entity: %v", err)
	}
	return err
}

func (r *repository[T]) Update(ctx context.Context, entity *T) error {
	r.log.Info("Updating an entity")

	err := r.db.WithContext(ctx).Save(entity).Error
	if err != nil {
		r.log.Errorf("Error updating entity: %v", err)
	}
	return err
}

func (r *repository[T]) Delete(ctx context.Context, id uint) error {
	r.log.Infof("Deleting entity with ID: %d", id)

	err := r.db.WithContext(ctx).Delete(new(T), id).Error
	if err != nil {
		r.log.Errorf("Error deleting entity with ID %d: %v", id, err)
	}
	return err
}

func (r *repository[T]) GetDB() *gorm.DB {
	return r.db
}
