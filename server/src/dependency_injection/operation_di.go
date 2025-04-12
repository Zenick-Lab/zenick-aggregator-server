package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewOperationUsecaseProvider() interfaces.IOperationUsecase {
	log := logrus.New()
	operationRepository := NewRepositoryProvider[model.Operation]()
	return usecase.NewOperationUsecase(operationRepository, log)
}
