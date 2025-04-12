package interfaces

import (
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"context"
)

type ITokenUsecase interface {
	GetAllTokens(ctx context.Context) ([]model.Token, error)
	GetTokenByID(ctx context.Context, id uint) (*model.Token, error)
	CreateToken(ctx context.Context, token *model.Token) error
	UpdateToken(ctx context.Context, token *model.Token) error
	DeleteToken(ctx context.Context, id uint) error
}
