package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"

	"github.com/sirupsen/logrus"
)

type historyUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.History]
}

func NewHistoryUsecase(repo interfaces.IRepository[model.History], log *logrus.Logger) interfaces.IHistoryUsecase {
	return &historyUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *historyUsecase) GetAllHistories(ctx context.Context) ([]model.History, error) {
	u.log.Info("Fetching all histories")
	return u.Repo.GetAll(ctx)
}

func (u *historyUsecase) GetHistoriesDetails(ctx context.Context) ([]dto.HistoryResponse, error) {
	var histories []model.History
	err := u.Repo.GetDB().WithContext(ctx).Preload("Provider").Preload("Token").Preload("Operation").Find(&histories).Error
	if err != nil {
		u.log.Errorf("Error fetching histories details: %v", err)
		return nil, err
	}

	var responses []dto.HistoryResponse
	for _, history := range histories {
		responses = append(responses, dto.HistoryResponse{
			Provider:  history.Provider.Name,
			Token:     history.Token.Name,
			Operation: history.Operation.Name,
			APR:       history.APR,
			CreatedAt: history.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

func (u *historyUsecase) GetHistoriesByCondition(ctx context.Context, req *dto.GetHistoryRequest) ([]dto.HistoryResponse, error) {
	var histories []model.History

	query := u.Repo.GetDB().WithContext(ctx).
		Preload("Provider").
		Preload("Token").
		Preload("Operation").
		Table("histories").
		Joins("JOIN providers ON providers.id = histories.provider_id").
		Joins("JOIN tokens ON tokens.id = histories.token_id").
		Joins("JOIN operations ON operations.id = histories.operation_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.Token != "" {
		query = query.Where("tokens.name ILIKE ?", "%"+req.Token+"%")
	}
	if req.Operation != "" {
		query = query.Where("operations.name ILIKE ?", "%"+req.Operation+"%")
	}
	if req.APR != nil {
		query = query.Where("histories.apr = ?", *req.APR)
	}
	if !req.FromDate.IsZero() {
		query = query.Where("created_at >= ?", req.FromDate)
	}

	if !req.ToDate.IsZero() {
		query = query.Where("created_at <= ?", req.ToDate)
	}
	err := query.Find(&histories).Error
	if err != nil {
		u.log.Errorf("Error fetching histories by condition: %v", err)
		return nil, err
	}

	var responses []dto.HistoryResponse
	for _, history := range histories {
		responses = append(responses, dto.HistoryResponse{
			Provider:  history.Provider.Name,
			Token:     history.Token.Name,
			Operation: history.Operation.Name,
			APR:       history.APR,
			CreatedAt: history.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

func (u *historyUsecase) GetHistoryByCondition(ctx context.Context, req *dto.GetNewestHistoryRequest) (*dto.HistoryResponse, error) {
	var history model.History

	query := u.Repo.GetDB().WithContext(ctx).
		Preload("Provider").
		Preload("Token").
		Preload("Operation").
		Table("histories").
		Joins("JOIN providers ON providers.id = histories.provider_id").
		Joins("JOIN tokens ON tokens.id = histories.token_id").
		Joins("JOIN operations ON operations.id = histories.operation_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.Token != "" {
		query = query.Where("tokens.name ILIKE ?", "%"+req.Token+"%")
	}
	if req.Operation != "" {
		query = query.Where("operations.name ILIKE ?", "%"+req.Operation+"%")
	}

	err := query.Order("created_at DESC").First(&history).Error
	if err != nil {
		u.log.Errorf("Error fetching newest history by condition: %v", err)
		return nil, err
	}

	response := &dto.HistoryResponse{
		Provider:  history.Provider.Name,
		Token:     history.Token.Name,
		Operation: history.Operation.Name,
		APR:       history.APR,
		CreatedAt: history.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return response, nil
}

func (u *historyUsecase) GetHistoryByID(ctx context.Context, id uint) (*model.History, error) {
	u.log.Infof("Fetching history with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *historyUsecase) CreateHistory(ctx context.Context, history *model.History) error {
	u.log.Info("Creating a new history")
	return u.Repo.Create(ctx, history)
}

func (u *historyUsecase) UpdateHistory(ctx context.Context, history *model.History) error {
	u.log.Infof("Updating history with ID: %d", history.ID)
	return u.Repo.Update(ctx, history)
}

func (u *historyUsecase) DeleteHistory(ctx context.Context, id uint) error {
	u.log.Infof("Deleting history with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
