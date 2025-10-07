package repo

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"context"
	"sync"
)

type memRepo struct {
	mutex        sync.RWMutex
	cards        map[string]*model.Card
	createdUsers map[string]*model.Card
}

func NewMemRepo(ctx context.Context) (IRepository, error) {
	return &memRepo{
		cards:        make(map[string]*model.Card),
		createdUsers: make(map[string]*model.Card),
	}, nil
}

func (repo *memRepo) Close() error {
	return nil
}

func (repo *memRepo) CreateCard(ctx context.Context, card *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[card.ID] = card
	repo.createdUsers[card.UserID] = card

	return nil
}

func (repo *memRepo) CountCardByUserID(ctx context.Context, userID string) (int32, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, ok := repo.createdUsers[userID]
	if ok {
		return 1, nil
	}
	return 0, nil
}

func (repo *memRepo) GetCardByID(ctx context.Context, id string) (*model.Card, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card, ok := repo.cards[id]
	if !ok {
		return nil, errmsg.CardNotFound
	}

	return card, nil
}

func (repo *memRepo) GetCardByUserID(ctx context.Context, userID string) (*model.Card, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card, ok := repo.createdUsers[userID]
	if !ok {
		return nil, errmsg.CardNotFound
	}

	return card, nil
}

func (repo *memRepo) UpdateCardStatus(ctx context.Context, id string, status model.Status) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[id].Status = status
	userID := repo.cards[id].UserID
	repo.createdUsers[userID].Status = status

	return nil
}
