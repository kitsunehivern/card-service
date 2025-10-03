package repo

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"sync"
)

type memRepo struct {
	mutex     sync.RWMutex
	cards     map[string]*model.Card
	userIndex map[string]string
}

func NewMemRepo() Repository {
	return &memRepo{
		cards:     make(map[string]*model.Card),
		userIndex: make(map[string]string),
	}
}

func (repo *memRepo) Create(card *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	if _, exists := repo.userIndex[card.UserID]; exists {
		return errmsg.CardAlreadyExists
	}

	repo.cards[card.ID] = card
	repo.userIndex[card.UserID] = card.ID

	return nil
}

func (repo *memRepo) Get(id string) (*model.Card, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card, ok := repo.cards[id]
	if !ok {
		return nil, errmsg.CardNotFound
	}

	return card, nil
}

func (repo *memRepo) Update(c *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[c.ID] = c

	return nil
}
