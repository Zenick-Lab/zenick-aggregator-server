package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"
	"time"

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

func (u *historyUsecase) GetHistoriesByCondition(ctx context.Context, req *dto.GetNewestHistoryRequest) ([]dto.HistoryResponse, error) {
	var results []struct {
		ProviderName  string
		TokenName     string
		OperationName string
		Link          string
		APR           float32
		CreatedAt     time.Time
	}

	subquery := u.Repo.GetDB().WithContext(ctx).
		Table("histories").
		Select("provider_id, MAX(created_at) AS latest_created_at").
		Group("provider_id")

	if req.Token != "" {
		subquery = subquery.Where("token_id IN (SELECT id FROM tokens WHERE name ILIKE ?)", "%"+req.Token+"%")
	}
	if req.Operation != "" {
		subquery = subquery.Where("operation_id IN (SELECT id FROM operations WHERE name ILIKE ?)", "%"+req.Operation+"%")
	}
	if req.Provider != "" {
		subquery = subquery.Where("provider_id IN (SELECT id FROM providers WHERE name ILIKE ?)", "%"+req.Provider+"%")
	}

	query := u.Repo.GetDB().WithContext(ctx).
		Table("histories").
		Select(`
            providers.name AS provider_name,
            tokens.name AS token_name,
            operations.name AS operation_name,
            history_links.link AS link,
            histories.apr AS apr,
            histories.created_at AS created_at
        `).
		Joins("JOIN providers ON providers.id = histories.provider_id").
		Joins("JOIN tokens ON tokens.id = histories.token_id").
		Joins("JOIN operations ON operations.id = histories.operation_id").
		Joins("LEFT JOIN history_links ON history_links.provider_id = histories.provider_id AND history_links.token_id = histories.token_id AND history_links.operation_id = histories.operation_id").
		Joins("INNER JOIN (?) AS latest_histories ON histories.provider_id = latest_histories.provider_id AND histories.created_at = latest_histories.latest_created_at", subquery)

	err := query.Scan(&results).Error
	if err != nil {
		u.log.Errorf("Error fetching histories by condition: %v", err)
		return nil, err
	}

	var responses []dto.HistoryResponse
	for _, result := range results {
		responses = append(responses, dto.HistoryResponse{
			Provider:  result.ProviderName,
			Token:     result.TokenName,
			Operation: result.OperationName,
			Link:      result.Link,
			APR:       result.APR,
			CreatedAt: result.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

func (u *historyUsecase) GetHistoryByCondition(ctx context.Context, req *dto.GetNewestHistoryRequest) (*dto.HistoryResponse, error) {
	var result struct {
		ProviderName  string
		TokenName     string
		OperationName string
		Link          string
		APR           float32
		CreatedAt     time.Time
	}

	query := u.Repo.GetDB().WithContext(ctx).
		Table("histories").
		Select(`
            providers.name AS provider_name,
            tokens.name AS token_name,
            operations.name AS operation_name,
            history_links.link AS link,
            histories.apr AS apr,
            histories.created_at AS created_at
        `).
		Joins("JOIN providers ON providers.id = histories.provider_id").
		Joins("JOIN tokens ON tokens.id = histories.token_id").
		Joins("JOIN operations ON operations.id = histories.operation_id").
		Joins("LEFT JOIN history_links ON history_links.provider_id = histories.provider_id AND history_links.token_id = histories.token_id AND history_links.operation_id = histories.operation_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.Token != "" {
		query = query.Where("tokens.name ILIKE ?", "%"+req.Token+"%")
	}
	if req.Operation != "" {
		query = query.Where("operations.name ILIKE ?", "%"+req.Operation+"%")
	}

	err := query.Order("histories.created_at DESC").Limit(1).Scan(&result).Error
	if err != nil {
		u.log.Errorf("Error fetching newest history by condition: %v", err)
		return nil, err
	}

	response := &dto.HistoryResponse{
		Provider:  result.ProviderName,
		Token:     result.TokenName,
		Operation: result.OperationName,
		Link:      result.Link,
		APR:       result.APR,
		CreatedAt: result.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
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
