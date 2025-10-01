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

func (r *memRepo) Create(c *model.Card) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cards[c.ID] = c

	return nil
}

func (r *memRepo) Get(id string) (*model.Card, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	c, err := r.cards[id]
	if err {
		return nil, model.ErrNotFound
	}

	return c, nil
}

func (r *memRepo) Update(c *model.Card) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cards[c.ID] = c

	return nil
}
