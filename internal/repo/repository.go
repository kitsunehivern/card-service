package repo

import (
	"card-service/internal/model"
	"context"
)

type IRepository interface {
	CreateCard(ctx context.Context, card *model.Card) error
	CountCardByUserID(ctx context.Context, userID string) (int, error)
	GetCardByID(ctx context.Context, id int64) (*model.Card, error)
	GetCardByUserID(ctx context.Context, userID string) (*model.Card, error)
	UpdateCardStatus(ctx context.Context, id int64, status model.Status) error
	CloseExpiredCard(ctx context.Context) error
}
