package usecase

import (
	"Zenick-Lab/zenick-aggregator-server/src/interfaces"
	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/model/dto"
	"context"

	"github.com/sirupsen/logrus"
)

type historyLinkUsecase struct {
	log  *logrus.Logger
	Repo interfaces.IRepository[model.HistoryLink]
}

func NewHistoryLinkUsecase(repo interfaces.IRepository[model.HistoryLink], log *logrus.Logger) interfaces.IHistoryLinkUsecase {
	return &historyLinkUsecase{
		log:  log,
		Repo: repo,
	}
}

func (u *historyLinkUsecase) GetAllHistoryLinks(ctx context.Context) ([]model.HistoryLink, error) {
	u.log.Info("Fetching all historyLinks")
	return u.Repo.GetAll(ctx)
}

func (u *historyLinkUsecase) GetHistoryLinksDetails(ctx context.Context) ([]dto.HistoryLinkResponse, error) {
	var historyLinks []model.HistoryLink
	err := u.Repo.GetDB().WithContext(ctx).Preload("Provider").Preload("Token").Preload("Operation").Find(&historyLinks).Error
	if err != nil {
		u.log.Errorf("Error fetching historyLinks details: %v", err)
		return nil, err
	}

	var responses []dto.HistoryLinkResponse
	for _, historyLink := range historyLinks {
		responses = append(responses, dto.HistoryLinkResponse{
			Provider:  historyLink.Provider.Name,
			Token:     historyLink.Token.Name,
			Operation: historyLink.Operation.Name,
			Link:      historyLink.Link,
		})
	}

	return responses, nil
}

func (u *historyLinkUsecase) GetHistoryLinkByCondition(ctx context.Context, req *dto.GetHistoryLinkRequest) (*dto.HistoryLinkResponse, error) {
	var historyLink model.HistoryLink

	query := u.Repo.GetDB().WithContext(ctx).
		Preload("Provider").
		Preload("Token").
		Preload("Operation").
		Table("history_links").
		Joins("JOIN providers ON providers.id = history_links.provider_id").
		Joins("JOIN tokens ON tokens.id = history_links.token_id").
		Joins("JOIN operations ON operations.id = history_links.operation_id")

	if req.Provider != "" {
		query = query.Where("providers.name ILIKE ?", "%"+req.Provider+"%")
	}
	if req.Token != "" {
		query = query.Where("tokens.name ILIKE ?", "%"+req.Token+"%")
	}
	if req.Operation != "" {
		query = query.Where("operations.name ILIKE ?", "%"+req.Operation+"%")
	}

	err := query.First(&historyLink).Error
	if err != nil {
		u.log.Errorf("Error fetching historyLink by condition: %v", err)
		return nil, err
	}

	response := &dto.HistoryLinkResponse{
		Provider:  historyLink.Provider.Name,
		Token:     historyLink.Token.Name,
		Operation: historyLink.Operation.Name,
		Link:      historyLink.Link,
	}

	return response, nil
}

func (u *historyLinkUsecase) GetHistoryLinkByID(ctx context.Context, id uint) (*model.HistoryLink, error) {
	u.log.Infof("Fetching historyLink with ID: %d", id)
	return u.Repo.GetByID(ctx, id)
}

func (u *historyLinkUsecase) CreateHistoryLink(ctx context.Context, historyLink *model.HistoryLink) error {
	u.log.Info("Creating a new historyLink")
	return u.Repo.Create(ctx, historyLink)
}

func (u *historyLinkUsecase) UpdateHistoryLink(ctx context.Context, historyLink *model.HistoryLink) error {
	u.log.Infof("Updating historyLink with ID: %d", historyLink.ID)
	return u.Repo.Update(ctx, historyLink)
}

func (u *historyLinkUsecase) DeleteHistoryLink(ctx context.Context, id uint) error {
	u.log.Infof("Deleting historyLink with ID: %d", id)
	return u.Repo.Delete(ctx, id)
}
