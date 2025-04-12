package dependency_injection

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/usecase"

	"github.com/sirupsen/logrus"
)

func NewTokenUsecaseProvider() interfaces.ITokenUsecase {
	log := logrus.New()
	tokenRepository := NewRepositoryProvider[model.Token]()
	return usecase.NewTokenUsecase(tokenRepository, log)
}
