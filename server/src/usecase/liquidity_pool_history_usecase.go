package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"

	"github.com/sirupsen/logrus"
)

type liquidityPoolHistoryUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.LiquidityPoolHistory]
}

func NewLiquidityPoolHistoryUsecase(repo interfaces.IRepository[model.LiquidityPoolHistory], log *logrus.Logger) interfaces.ILiquidityPoolHistoryUsecase {
	return &liquidityPoolHistoryUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *liquidityPoolHistoryUsecase) GetAllLiquidityPoolHistories(ctx context.Context) ([]model.LiquidityPoolHistory, error) {
	u.log.Info("Fetching all Liquidity Pool Histories")
	return u.Repo.GetAll(ctx)
}

func (u *liquidityPoolHistoryUsecase) GetLiquidityPoolHistoriesDetails(ctx context.Context) ([]dto.LiquidityPoolHistoryResponse, error) {
	var liquidityPoolHistories []model.LiquidityPoolHistory
	err := u.Repo.GetDB().WithContext(ctx).Preload("Provider").Preload("TokenA").Preload("TokenB").Find(&liquidityPoolHistories).Error
	if err != nil {
		u.log.Errorf("Error fetching Liquidity Pool Histories details: %v", err)
		return nil, err
	}

	var responses []dto.LiquidityPoolHistoryResponse
	for _, liquidityPoolHistory := range liquidityPoolHistories {
		responses = append(responses, dto.LiquidityPoolHistoryResponse{
			Provider:  liquidityPoolHistory.Provider.Name,
			TokenA:    liquidityPoolHistory.TokenA.Name,
			TokenB:    liquidityPoolHistory.TokenB.Name,
			Link:      liquidityPoolHistory.Link,
			APR:       liquidityPoolHistory.APR,
			CreatedAt: liquidityPoolHistory.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

func (u *liquidityPoolHistoryUsecase) GetLiquidityPoolHistoryByCondition(ctx context.Context, req *dto.GetNewestLiquidityPoolHistoryRequest) (*dto.LiquidityPoolHistoryResponse, error) {
	var history model.LiquidityPoolHistory

	query := u.Repo.GetDB().WithContext(ctx).
		Preload("Provider").
		Preload("TokenA").
		Preload("TokenB").
		Table("liquidity_pool_histories").
		Joins("JOIN providers ON providers.id = liquidity_pool_histories.provider_id").
		Joins("JOIN tokens AS token_a ON token_a.id = liquidity_pool_histories.token_a_id").
		Joins("JOIN tokens AS token_b ON token_b.id = liquidity_pool_histories.token_b_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.TokenA != "" {
		query = query.Where("token_a.name ILIKE ?", "%"+req.TokenA+"%")
	}
	if req.TokenB != "" {
		query = query.Where("token_b.name ILIKE ?", "%"+req.TokenB+"%")
	}

	err := query.Order("created_at DESC").First(&history).Error
	if err != nil {
		u.log.Errorf("Error fetching newest liquidity pool history by condition: %v", err)
		return nil, err
	}

	response := &dto.LiquidityPoolHistoryResponse{
		Provider:  history.Provider.Name,
		TokenA:    history.TokenA.Name,
		TokenB:    history.TokenB.Name,
		Link:      history.Link,
		APR:       history.APR,
		CreatedAt: history.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return response, nil
}

func (u *liquidityPoolHistoryUsecase) GetLiquidityPoolHistoryByID(ctx context.Context, id uint) (*model.LiquidityPoolHistory, error) {
	u.log.Infof("Fetching Liquidity Pool History with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *liquidityPoolHistoryUsecase) CreateLiquidityPoolHistory(ctx context.Context, liquidityPoolHistory *model.LiquidityPoolHistory) error {
	u.log.Info("Creating a new Liquidity Pool History")
	return u.Repo.Create(ctx, liquidityPoolHistory)
}

func (u *liquidityPoolHistoryUsecase) UpdateLiquidityPoolHistory(ctx context.Context, liquidityPoolHistory *model.LiquidityPoolHistory) error {
	u.log.Infof("Updating Liquidity Pool History with ID: %d", liquidityPoolHistory.ID)
	return u.Repo.Update(ctx, liquidityPoolHistory)
}

func (u *liquidityPoolHistoryUsecase) DeleteLiquidityPoolHistory(ctx context.Context, id uint) error {
	u.log.Infof("Deleting Liquidity Pool History with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
