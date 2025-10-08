package repo

import (
	"card-service/internal/model"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type psqlRepo struct {
	db *gorm.DB
}

func NewPsqlRepo(ctx context.Context, dsn string) (IRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &psqlRepo{db: db}, nil
}

func (repo *psqlRepo) CreateCard(ctx context.Context, card *model.Card) error {
	if result := repo.db.WithContext(ctx).Create(card); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *psqlRepo) CountCardByUserID(ctx context.Context, userID string) (int, error) {
	var count int64
	repo.db.WithContext(ctx).Model(&model.Card{}).Where("user_id = ?", userID).Count(&count)

	return int(count), nil
}

func (repo *psqlRepo) GetCardByID(ctx context.Context, id string) (*model.Card, error) {
	var card model.Card
	if result := repo.db.WithContext(ctx).Where("id = ?", id).First(&card); result.Error != nil {
		return nil, result.Error
	}

	return &card, nil
}

func (repo *psqlRepo) GetCardByUserID(ctx context.Context, userID string) (*model.Card, error) {
	var card model.Card
	if result := repo.db.WithContext(ctx).Where("user_id = ?", userID).First(&card); result.Error != nil {
		return nil, result.Error
	}

	return &card, nil
}

func (repo *psqlRepo) UpdateCardStatus(ctx context.Context, id string, status model.Status) error {
	repo.db.WithContext(ctx).Model(&model.Card{}).Where("id = ?", id).Update("status", status)
	return nil
}
