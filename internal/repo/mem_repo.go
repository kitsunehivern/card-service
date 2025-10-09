package repo

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"context"
	"sync"
)

type memRepo struct {
	mutex        sync.RWMutex
	lastID       int64
	cards        map[int64]*model.Card
	createdUsers map[string]*model.Card
}

func NewMemRepo(ctx context.Context) (IRepository, error) {
	return &memRepo{
		lastID:       0,
		cards:        make(map[int64]*model.Card),
		createdUsers: make(map[string]*model.Card),
	}, nil
}

func (repo *memRepo) CreateCard(ctx context.Context, card *model.Card) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	card.ID = repo.lastID + 1
	repo.lastID++
	repo.cards[card.ID] = card
	repo.createdUsers[card.UserID] = card

	return nil
}

func (repo *memRepo) CountCardByUserID(ctx context.Context, userID string) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, ok := repo.createdUsers[userID]
	if ok {
		return 1, nil
	}
	return 0, nil
}

func (repo *memRepo) GetCardByID(ctx context.Context, id int64) (*model.Card, error) {
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

func (repo *memRepo) UpdateCardStatus(ctx context.Context, id int64, status model.Status) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.cards[id].Status = status
	userID := repo.cards[id].UserID
	repo.createdUsers[userID].Status = status

	return nil
}
