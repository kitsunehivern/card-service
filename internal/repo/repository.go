package repo

import "card-service/internal/model"

type IRepository interface {
	CreateCard(card *model.Card) error
	CountCard(userID string) (int32, error)
	GetCard(ID string) (*model.Card, error)
	UpdateCard(card *model.Card) error
}
