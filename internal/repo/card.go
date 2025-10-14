package repo

import (
	"card-service/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type ICardRepo interface {
	CreateCard(ctx context.Context, card *model.Card) error
	CountCard(ctx context.Context, params model.CardParams) (int64, error)
	GetCard(ctx context.Context, params model.CardParams) (*model.Card, error)
	UpdateCardStatus(ctx context.Context, params model.CardParams, status model.CardStatus) error
	CloseExpiredCard(ctx context.Context) error
}

type CardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) ICardRepo {
	return &CardRepo{db: db}
}

func (repo *CardRepo) buildCardQuery(ctx context.Context, params model.CardParams) model.CardQuerySet {
	query := model.NewCardQuerySet(repo.db.WithContext(ctx))
	if params.ID > 0 {
		query = query.IDEq(params.ID)
	}

	if params.UserID > 0 {
		query = query.UserIDEq(params.UserID)
	}

	return query
}

func (repo *CardRepo) CreateCard(ctx context.Context, card *model.Card) error {
	return repo.db.WithContext(ctx).Create(card).Error
}

func (repo *CardRepo) CountCard(ctx context.Context, params model.CardParams) (int64, error) {
	count, err := repo.buildCardQuery(ctx, params).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *CardRepo) GetCard(ctx context.Context, params model.CardParams) (*model.Card, error) {
	var card model.Card
	if err := repo.buildCardQuery(ctx, params).One(&card); err != nil {
		return nil, err
	}

	return &card, nil
}

func (repo *CardRepo) UpdateCardStatus(ctx context.Context, params model.CardParams, status model.CardStatus) error {
	return repo.buildCardQuery(ctx, params).GetDB().Update("status", status).Error
}

func (repo *CardRepo) CloseExpiredCard(ctx context.Context) error {
	return repo.db.
		WithContext(ctx).
		Model(&model.Card{}).
		Where("expiration_date < ?", time.Now().UTC()).Where("status != ?", model.CardStatusClosed).
		Update("status", model.CardStatusClosed).
		Error
}
