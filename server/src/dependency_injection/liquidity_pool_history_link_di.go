package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewLiquidityPoolHistoryLinkUsecaseProvider() interfaces.ILiquidityPoolHistoryLinkUsecase {
	log := logrus.New()
	historyRepository := NewRepositoryProvider[model.LiquidityPoolHistoryLink]()
	return usecase.NewLiquidityPoolHistoryLinkUsecase(historyRepository, log)
}
