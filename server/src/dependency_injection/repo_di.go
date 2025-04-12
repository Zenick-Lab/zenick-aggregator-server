package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/pkg/postgresql"
	"Zenick-Lab/zenick-aggregator-server/src/repository"

	"github.com/sirupsen/logrus"
)

func NewRepositoryProvider[T any]() interfaces.IRepository[T] {
	log := logrus.New()

	db, err := postgresql.NewGormDB()
	if err != nil {
		log.Error(err)
	}

	return repository.NewRepository[T](db, log)
}
