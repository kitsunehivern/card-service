package repo

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"sync"
)

type memRepo struct {
	mutex        sync.RWMutex
	cards        map[string]*model.Card
	createdUsers map[string]bool
}

func NewMemRepo() IRepository {
	return &memRepo{
		cards:        make(map[string]*model.Card),
		createdUsers: make(map[string]bool),
	}
}

func (repo *memRepo) CreateCard(card *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[card.ID] = card
	repo.createdUsers[card.UserID] = true

	return nil
}

func (repo *memRepo) CountCard(userID string) (int32, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, ok := repo.createdUsers[userID]
	if ok {
		return 1, nil
	}
	return 0, nil
}

func (repo *memRepo) GetCard(ID string) (*model.Card, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card, ok := repo.cards[ID]
	if !ok {
		return nil, errmsg.CardNotFound
	}

	return card, nil
}

func (repo *memRepo) UpdateCard(c *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[c.ID] = c

	return nil
}
