package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"

	"github.com/sirupsen/logrus"
)

type liquidityPoolHistoryLinkUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.LiquidityPoolHistoryLink]
}

func NewLiquidityPoolHistoryLinkUsecase(repo interfaces.IRepository[model.LiquidityPoolHistoryLink], log *logrus.Logger) interfaces.ILiquidityPoolHistoryLinkUsecase {
	return &liquidityPoolHistoryLinkUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *liquidityPoolHistoryLinkUsecase) GetAllLiquidityPoolHistoryLinks(ctx context.Context) ([]model.LiquidityPoolHistoryLink, error) {
	u.log.Info("Fetching all LiquidityPoolHistoryLinks")
	return u.Repo.GetAll(ctx)
}

func (u *liquidityPoolHistoryLinkUsecase) GetLiquidityPoolHistoryLinksDetails(ctx context.Context) ([]dto.LiquidityPoolHistoryLinkResponse, error) {
	var liquidityPoolHistoryLinks []model.LiquidityPoolHistoryLink
	err := u.Repo.GetDB().WithContext(ctx).Preload("Provider").Preload("TokenA").Preload("TokenB").Find(&liquidityPoolHistoryLinks).Error
	if err != nil {
		u.log.Errorf("Error fetching LiquidityPoolHistoryLinks details: %v", err)
		return nil, err
	}

	var responses []dto.LiquidityPoolHistoryLinkResponse
	for _, liquidityPoolHistoryLink := range liquidityPoolHistoryLinks {
		responses = append(responses, dto.LiquidityPoolHistoryLinkResponse{
			Provider: liquidityPoolHistoryLink.Provider.Name,
			TokenA:   liquidityPoolHistoryLink.TokenA.Name,
			TokenB:   liquidityPoolHistoryLink.TokenB.Name,
			Link:     liquidityPoolHistoryLink.Link,
		})
	}

	return responses, nil
}

func (u *liquidityPoolHistoryLinkUsecase) GetLiquidityPoolHistoryLinkByCondition(ctx context.Context, req *dto.GetLiquidityPoolHistoryLinkRequest) (*dto.LiquidityPoolHistoryLinkResponse, error) {
	var history model.LiquidityPoolHistoryLink

	query := u.Repo.GetDB().WithContext(ctx).
		Preload("Provider").
		Preload("TokenA").
		Preload("TokenB").
		Table("liquidity_pool_history_links").
		Joins("JOIN providers ON providers.id = liquidity_pool_history_links.provider_id").
		Joins("JOIN tokens AS token_a ON token_a.id = liquidity_pool_history_links.token_a_id").
		Joins("JOIN tokens AS token_b ON token_b.id = liquidity_pool_history_links.token_b_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.TokenA != "" {
		query = query.Where("token_a.name ILIKE ?", "%"+req.TokenA+"%")
	}
	if req.TokenB != "" {
		query = query.Where("token_b.name ILIKE ?", "%"+req.TokenB+"%")
	}

	err := query.First(&history).Error
	if err != nil {
		u.log.Errorf("Error fetching liquidity pool history link by condition: %v", err)
		return nil, err
	}

	response := &dto.LiquidityPoolHistoryLinkResponse{
		Provider: history.Provider.Name,
		TokenA:   history.TokenA.Name,
		TokenB:   history.TokenB.Name,
		Link:     history.Link,
	}

	return response, nil
}

func (u *liquidityPoolHistoryLinkUsecase) GetLiquidityPoolHistoryLinkByID(ctx context.Context, id uint) (*model.LiquidityPoolHistoryLink, error) {
	u.log.Infof("Fetching LiquidityPoolHistoryLink with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *liquidityPoolHistoryLinkUsecase) CreateLiquidityPoolHistoryLink(ctx context.Context, liquidityPoolHistoryLink *model.LiquidityPoolHistoryLink) error {
	u.log.Info("Creating a new LiquidityPoolHistoryLink")
	return u.Repo.Create(ctx, liquidityPoolHistoryLink)
}

func (u *liquidityPoolHistoryLinkUsecase) UpdateLiquidityPoolHistoryLink(ctx context.Context, liquidityPoolHistoryLink *model.LiquidityPoolHistoryLink) error {
	u.log.Infof("Updating LiquidityPoolHistoryLink with ID: %d", liquidityPoolHistoryLink.ID)
	return u.Repo.Update(ctx, liquidityPoolHistoryLink)
}

func (u *liquidityPoolHistoryLinkUsecase) DeleteLiquidityPoolHistoryLink(ctx context.Context, id uint) error {
	u.log.Infof("Deleting LiquidityPoolHistoryLink with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
