package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewHistoryLinkUsecaseProvider() interfaces.IHistoryLinkUsecase {
	log := logrus.New()
	historyLinkRepository := NewRepositoryProvider[model.HistoryLink]()
	return usecase.NewHistoryLinkUsecase(historyLinkRepository, log)
}
