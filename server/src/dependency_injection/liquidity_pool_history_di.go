package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewLiquidityPoolHistoryUsecaseProvider() interfaces.ILiquidityPoolHistoryUsecase {
	log := logrus.New()
	historyRepository := NewRepositoryProvider[model.LiquidityPoolHistory]()
	return usecase.NewLiquidityPoolHistoryUsecase(historyRepository, log)
}
