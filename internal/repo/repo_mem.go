package repo

import (
	"card-service/internal/model"
	"sync"
)

type memRepo struct {
	cards map[string]*model.Card
	mutex sync.RWMutex
}

func NewMemRepo() Repository {
	return &memRepo{
		cards: map[string]*model.Card{},
	}
}

func (repo *memRepo) Create(card *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[card.ID] = card

	return nil
}

func (repo *memRepo) Get(id string) (*model.Card, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card, err := repo.cards[id]
	if err {
		return nil, model.ErrNotFound
	}

	return card, nil
}

func (repo *memRepo) Update(c *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[c.ID] = c

	return nil
}
