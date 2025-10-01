package repo

import "card-service/internal/model"

type Repository interface {
	Create(*model.Card) error
	Get(id string) (*model.Card, error)
	Update(*model.Card) error
}
