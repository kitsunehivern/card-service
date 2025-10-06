package repo

import "card-service/internal/model"

type IRepository interface {
	CreateCard(*model.Card) error
	HasCreatedCard(string) (bool, error)
	GetCard(string) (*model.Card, error)
	UpdateCard(*model.Card) error
}
