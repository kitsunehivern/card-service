package repo

import (
	"card-service/internal/model"
	"context"
)

type IRepository interface {
	CreateCard(ctx context.Context, card *model.Card) error
	CountCardByUserID(ctx context.Context, userID string) (int, error)
	GetCardByID(ctx context.Context, id string) (*model.Card, error)
	GetCardByUserID(ctx context.Context, userID string) (*model.Card, error)
	UpdateCardStatus(ctx context.Context, id string, status model.Status) error
}
