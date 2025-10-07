package repo

import "card-service/internal/model"

type IRepository interface {
	CreateCard(*model.Card) error
	CountCard(string) (int32, error)
	GetCard(string) (*model.Card, error)
	UpdateCard(*model.Card) error
}
