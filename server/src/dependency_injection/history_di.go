package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewHistoryUsecaseProvider() interfaces.IHistoryUsecase {
	log := logrus.New()
	historyRepository := NewRepositoryProvider[model.History]()
	return usecase.NewHistoryUsecase(historyRepository, log)
}
