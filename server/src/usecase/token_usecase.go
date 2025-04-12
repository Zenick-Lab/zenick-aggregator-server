package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"

	"github.com/sirupsen/logrus"
)

type tokenUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.Token]
}

func NewTokenUsecase(repo interfaces.IRepository[model.Token], log *logrus.Logger) interfaces.ITokenUsecase {
	return &tokenUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *tokenUsecase) GetAllTokens(ctx context.Context) ([]model.Token, error) {
	u.log.Info("Fetching all tokens")
	return u.Repo.GetAll(ctx)
}

func (u *tokenUsecase) GetTokenByID(ctx context.Context, id uint) (*model.Token, error) {
	u.log.Infof("Fetching token with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *tokenUsecase) CreateToken(ctx context.Context, token *model.Token) error {
	u.log.Info("Creating a new token")
	return u.Repo.Create(ctx, token)
}

func (u *tokenUsecase) UpdateToken(ctx context.Context, token *model.Token) error {
	u.log.Infof("Updating token with ID: %d", token.ID)
	return u.Repo.Update(ctx, token)
}

func (u *tokenUsecase) DeleteToken(ctx context.Context, id uint) error {
	u.log.Infof("Deleting token with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
